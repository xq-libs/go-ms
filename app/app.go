package app

import (
	"github.com/xq-libs/go-ms/config"
	"log"
	"net/http"
	"time"
)

const (
	appCfgSectionName    = "app"
	serverCfgSectionName = "server"
)

var (
	appCfg    *AppConfig
	serverCfg *ServerConfig
)

func init() {
	log.Printf("Load app config data...")
	// 1.Load app config
	appCfg = new(AppConfig)
	config.GetDecryptSectionData(appCfgSectionName, appCfg)

	// 2.Load server config
	serverCfg = new(ServerConfig)
	config.GetSectionData(serverCfgSectionName, serverCfg)

	// 3.
	log.Println("Load app config data done")
}

func Start(h http.Handler) {
	// Listen and Server in 0.0.0.0:8080
	s := &http.Server{
		Addr:           serverCfg.Host + ":" + serverCfg.Port,
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	// Start Server
	log.Printf("App Server started at: %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil {
		log.Fatalf("App Server stop with error %v", err)
	}
}
