package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/xq-libs/go-utils/stringutil"
	"log"
	"net/http"
	"strconv"
)

// -----------------------------------------------------
// Get data from request
//

func GetRequestUser(c *gin.Context) User {
	return User{
		ID:       "100001",
		TenantId: "100001",
		Account:  "Admin",
		Name:     "超级管理员",
	}
}

func GetRequestPage(c *gin.Context) Pageable {
	return Pageable{
		Page: GetQueryInt(c, "page", 0),
		Size: GetQueryInt(c, "size", 10),
	}
}

func GetQueryInt(c *gin.Context, key string, df int) int {
	v := c.Query(key)
	if stringutil.IsNotBlank(v) {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Panicf("Value %s is not convert to int", v)
		}
		return i
	}
	return df
}

func GetParamInt(c *gin.Context, key string, df int) int {
	v := c.Param(key)
	if stringutil.IsNotBlank(v) {
		i, err := strconv.Atoi(v)
		if err != nil {
			log.Panicf("Value %s is not convert to int", v)
		}
		return i
	}
	return df
}

func GetRequestBody[T any](c *gin.Context, t T) T {
	err := c.ShouldBind(&t)
	if err != nil {
		log.Panicf("Bind obj from body failure: %v", err)
	}
	return t
}

func GetRequestQuery[T any](c *gin.Context, t T) T {
	err := c.ShouldBindQuery(&t)
	if err != nil {
		log.Panicf("Bind obj from query failure: %v", err)
	}
	return t
}

// -----------------------------------------------------
// Write data to response
//

type VoidResponseHandler func(ctx *gin.Context) error
type DataResponseHandler[T any] func(ctx *gin.Context) (T, error)

func HandleVoidResponse(handle VoidResponseHandler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		err := handle(ctx)
		SetResponse(ctx, "", err)
	}
}
func HandleDataResponse[T any](handle DataResponseHandler[T]) func(ctx *gin.Context) {
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
