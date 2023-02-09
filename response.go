package ms

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response[T any] struct {
	Code int64  `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data,omitempty"`
}

func NewResponse[T any](data T, err error) Response[T] {
	if err != nil {
		return Response[T]{
			Code: -1,
			Msg:  err.Error(),
		}
	}
	return Response[T]{
		Code: 0,
		Msg:  "Ok",
		Data: data,
	}
}

func ResponseSuccess() Response[any] {
	return Response[any]{
		Code: SuccessCode,
		Msg:  Success.DefaultValue,
		Data: nil,
	}
}

func ResponseError(err error) Response[any] {
	return Response[any]{
		Code: ServerErrorCode,
		Msg:  err.Error(),
		Data: nil,
	}
}

func ResponseWithData[T any](data T) Response[T] {
	return Response[T]{
		Code: SuccessCode,
		Msg:  Success.DefaultValue,
		Data: data,
	}
}

func ResponseWithError(err Error) Response[any] {
	return Response[any]{
		Code: err.Code(),
		Msg:  err.Error(),
		Data: nil,
	}
}

type VoidResponseHandler func(ctx *gin.Context) error

type DataResponseHandler[T any] func(ctx *gin.Context) (T, error)

func HandleResponseVoid(handle VoidResponseHandler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := handle(ctx)
		SetResponse(ctx, "", err)
	}
}

func HandleResponseData[T any](handle DataResponseHandler[T]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		data, err := handle(ctx)
		SetResponse(ctx, data, err)
	}
}

func SetResponse[T any](ctx *gin.Context, data T, err error) {
	if err != nil {
		switch err.(type) {
		case Error:
			ctx.JSON(http.StatusOK, ResponseWithError(err.(Error)))
		case error:
			ctx.JSON(http.StatusInternalServerError, ResponseWithError(NewError(err, ServerError)))
		default:
			ctx.JSON(http.StatusInternalServerError, ResponseWithError(NewError(nil, UnknownError)))
		}
	} else {
		ctx.JSON(http.StatusOK, ResponseWithData(data))
	}
}
