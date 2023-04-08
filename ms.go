package ms

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xq-libs/go-ms/config"
	"github.com/xq-libs/go-ms/database"
	"github.com/xq-libs/go-ms/server"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Listener listener interface
type Listener interface {
	OnStart(s *http.Server)
	OnShutdown()
}

var (
	lcx       = sync.Mutex{}
	listeners = make([]Listener, 0)
)

// StartServer Start server with http handler
func StartServer(h http.Handler) {
	// Listen and Server in 0.0.0.0:8080
	log.Println("App server will create with config data.")
	s := server.NewServer(h)

	// Register start hook
	s.RegisterOnShutdown(onShutdown)

	// Invoke on start
	onStart(s)

	// Start Server
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("App Server stop with error %v \n", err)
		}
	}()
	log.Printf("App Server started at: %s \n", s.Addr)
	waitExit(s)
}

func onStart(s *http.Server) {
	for _, l := range listeners {
		l.OnStart(s)
	}
}

func onShutdown() {
	for _, l := range listeners {
		l.OnShutdown()
	}
}

func waitExit(s *http.Server) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, os.Kill, syscall.SIGHUP, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-ch
	log.Printf("App server got a exit signal: %v", sig)
	// Create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Shutdown
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("App server shutdown failure: %v", err)
	}
	log.Println("App exist success.")
}

// AddListener add listener
func AddListener(l Listener) {
	lcx.Lock()
	defer lcx.Unlock()

	listeners = append(listeners, l)
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
