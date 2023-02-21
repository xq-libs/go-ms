package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xq-libs/go-ms/locale"
)

func NewLocalizer(ctx *gin.Context) *i18n.Localizer {
	acceptLang := ctx.GetHeader("Accept-Language")
	return locale.NewLocalizer(acceptLang)
}

func LocalizeMessage(ctx *gin.Context, message Message) string {
	localizer := NewLocalizer(ctx)
	return localizer.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    message.Key,
			Other: message.DefaultValue,
		},
		TemplateData: message.ArgMap,
	})
}
