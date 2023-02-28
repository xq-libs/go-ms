package ms

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xq-libs/go-ms/config"
	"github.com/xq-libs/go-ms/database"
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

func CustomerRecoveryFunc(sourceCtx *gin.Context, err any) {
	ctx := NewContext(sourceCtx)
	switch e := err.(type) {
	case Error:
		ctx.ResponseJson(http.StatusOK, ctx.GetErrorResponse(e))
	case error:
		ctx.ResponseJson(http.StatusInternalServerError, ctx.GetErrorResponse(NewError2(e, ServerError)))
	case string:
		ctx.ResponseJson(http.StatusInternalServerError, ctx.GetErrorResponse(NewError2(errors.New(e), ServerError)))
	default:
		str := fmt.Sprintf("%v", err)
		ctx.ResponseJson(http.StatusInternalServerError, ctx.GetErrorResponse(NewError2(errors.New(str), UnknownError)))
	}
}
