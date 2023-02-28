package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xq-libs/go-ms/locale"
	"github.com/xq-libs/go-utils/stringutil"
	"net/http"
)

// -----------------------------------------------------
// Get data from request
//

type Context struct {
	SourceCtx *gin.Context
}

func NewContext(ctx *gin.Context) *Context {
	return &Context{
		SourceCtx: ctx,
	}
}

// GetSourceContext Get original context
func (ctx *Context) GetSourceContext() *gin.Context {
	return ctx.SourceCtx
}

func (ctx *Context) GetRequestUser() User {
	return User{
		ID:       "100001",
		TenantId: "100001",
		Account:  "Admin",
		Name:     "超级管理员",
	}
}

func (ctx *Context) GetRequestPage() Pageable {
	return Pageable{
		Page: ctx.GetQueryInt("page", 0),
		Size: ctx.GetQueryInt("size", 10),
	}
}

// MustGetQueryInt Must get int from request query
func (ctx *Context) MustGetQueryInt(key string) int {
	v := ctx.SourceCtx.Query(key)
	if stringutil.IsBlank(v) {
		panic(NewError(RequestQueryError.AppendParam("param", key)))
	}
	return stringutil.ToInt(v)
}
func (ctx *Context) GetQueryInt(key string, df int) int {
	v := ctx.SourceCtx.Query(key)
	if stringutil.IsNotBlank(v) {
		return stringutil.ToInt(v)
	}
	return df
}

// MustGetQueryInt64 Must get int64 from request query
func (ctx *Context) MustGetQueryInt64(key string) int64 {
	v := ctx.SourceCtx.Query(key)
	if stringutil.IsBlank(v) {
		panic(NewError(RequestQueryError.AppendParam("param", key)))
	}
	return stringutil.ToInt64(v)
}
func (ctx *Context) GetQueryInt64(key string, df int64) int64 {
	v := ctx.SourceCtx.Query(key)
	if stringutil.IsNotBlank(v) {
		return stringutil.ToInt64(v)
	}
	return df
}

// GetParamInt Get int from path
func (ctx *Context) GetParamInt(key string, df int) int {
	v := ctx.SourceCtx.Param(key)
	if stringutil.IsNotBlank(v) {
		return stringutil.ToInt(v)
	}
	return df
}
func (ctx *Context) GetParamInt64(key string, df int64) int64 {
	v := ctx.SourceCtx.Param(key)
	if stringutil.IsNotBlank(v) {
		return stringutil.ToInt64(v)
	}
	return df
}

// MustGetQuery Must get string from request query
func (ctx *Context) MustGetQuery(key string) string {
	v := ctx.SourceCtx.Query(key)
	if stringutil.IsBlank(v) {
		panic(NewError(RequestQueryError.AppendParam("param", key)))
	}
	return v
}
func (ctx *Context) GetQuery(key string, defaultValue string) string {
	v := ctx.SourceCtx.Query(key)
	if stringutil.IsNotBlank(v) {
		return v
	}
	return defaultValue
}

// MustGetParam Get string from request path
func (ctx *Context) MustGetParam(key string) string {
	v := ctx.SourceCtx.Param(key)
	if stringutil.IsBlank(v) {
		panic(NewError(RequestParamError.AppendParam("param", key)))
	}
	return v
}
func (ctx *Context) GetParam(key string, defaultValue string) string {
	v := ctx.SourceCtx.Param(key)
	if stringutil.IsNotBlank(v) {
		return v
	}
	return defaultValue
}

// MustGetRequestBody Get request body
func (ctx *Context) MustGetRequestBody(obj any) {
	err := ctx.SourceCtx.ShouldBind(obj)
	if err != nil {
		panic(NewError2(err, RequestBodyBindError))
	}
}

// MustGetRequestJsonBody Get request body by json
func (ctx *Context) MustGetRequestJsonBody(obj any) {
	err := ctx.SourceCtx.ShouldBindJSON(obj)
	if err != nil {
		panic(NewError2(err, RequestBodyBindError))
	}
}

// MustGetRequestQuery Get request body by json
func (ctx *Context) MustGetRequestQuery(obj any) {
	err := ctx.SourceCtx.ShouldBindQuery(obj)
	if err != nil {
		panic(NewError2(err, RequestQueryBindError))
	}
}

func (ctx *Context) GetErrorResponse(err Error) Response[any] {
	return Response[any]{
		Code: err.Code(),
		Msg:  ctx.LocalizeMessage(err.Message),
	}
}

func (ctx *Context) GetSuccessResponse(data any) Response[any] {
	return Response[any]{
		Code: SuccessCode,
		Msg:  ctx.LocalizeMessage(Success),
		Data: data,
	}
}

func (ctx *Context) Response(data any, err error) {
	if err != nil {
		switch err.(type) {
		case Error:
			ctx.ResponseJson(http.StatusOK, ctx.GetErrorResponse(err.(Error)))
		case error:
			ctx.ResponseJson(http.StatusInternalServerError, ctx.GetErrorResponse(NewError2(err, ServerError)))
		default:
			ctx.ResponseJson(http.StatusInternalServerError, ctx.GetErrorResponse(NewError(UnknownError)))
		}
	} else {
		ctx.ResponseJson(http.StatusOK, ctx.GetSuccessResponse(data))
	}
}

func (ctx *Context) ResponseJson(code int, obj any) {
	ctx.SourceCtx.JSON(code, obj)
}

// GetLocalizer Create a Localizer for acquire localize message
func (ctx *Context) GetLocalizer() *i18n.Localizer {
	acceptLang := ctx.SourceCtx.GetHeader("Accept-Language")
	return locale.NewLocalizer(acceptLang)
}

// LocalizeMessage Localize message with Localizer
func (ctx *Context) LocalizeMessage(message Message) string {
	localizer := ctx.GetLocalizer()
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    message.Key,
			Other: message.DefaultValue,
		},
		TemplateData: message.ParamMap,
	})
}

// -----------------------------------------------------------
// Common get request method

func MustGetRequestBody[T any](c *Context, t T) T {
	c.MustGetRequestBody(t)
	return t
}
func MustGetRequestJsonBody[T any](c *Context, t T) T {
	c.MustGetRequestJsonBody(t)
	return t
}
func MustGetRequestQuery[T any](c *Context, t T) T {
	c.MustGetRequestQuery(t)
	return t
}

// -----------------------------------------------------
// Write data to response
//

type VoidResponseHandler func(ctx *Context) error
type DataResponseHandler[T any] func(ctx *Context) (T, error)

func HandleVoidResponse(handle VoidResponseHandler) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		c.Response("", handle(c))
	}
}

func HandleDataResponse[T any](handle DataResponseHandler[T]) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		c := NewContext(ctx)
		c.Response(handle(c))
	}
}
