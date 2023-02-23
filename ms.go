package ms

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xq-libs/go-ms/config"
	"github.com/xq-libs/go-ms/database"
	"github.com/xq-libs/go-ms/locale"
	"github.com/xq-libs/go-ms/server"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// StartServer Start server with http handler
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

// GetConfigData Get config data with section name
func GetConfigData(name string, data interface{}) {
	config.GetSectionData(name, data)
}

// GetDecryptConfigData Get config data with section name and decrypt it by jasypt
func GetDecryptConfigData(name string, data interface{}) {
	config.GetDecryptSectionData(name, data)
}

// NewLocalizer Create a Localizer for acquire localize message
func NewLocalizer(ctx *gin.Context) *i18n.Localizer {
	acceptLang := ctx.GetHeader("Accept-Language")
	return locale.NewLocalizer(acceptLang)
}

// LocalizeMessage Localize message with Localizer
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

// GetDb get db instance
func GetDb() *gorm.DB {
	return database.GetDb()
}

// WrapDbError wrap db error
func WrapDbError(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Error{
			Cause:   err,
			Message: DBNotFoundError,
		}
	} else if err != nil {
		return Error{
			Cause:   err,
			Message: DBError,
		}
	} else {
		return nil
	}
}
