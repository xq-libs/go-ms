package ms

import (
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xq-libs/go-ms/locale"
	"github.com/xq-libs/go-ms/server"
	"log"
	"net/http"
)

func StartServer(h http.Handler) {
	// Listen and Server in 0.0.0.0:8080
	s := server.NewServer(h)

	// Start Server
	log.Printf("App Server started at: %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("App Server stop with error %v", err)
	}
}

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
