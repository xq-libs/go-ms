package server

import (
	"github.com/xq-libs/go-ms/config"
	"log"
	"net/http"
	"time"
)

const (
	serverCfgSectionName = "server"
)

var (
	cfg *Config
)

func init() {
	log.Printf("Load server config data...")

	// 1.Load server config
	cfg = new(Config)
	config.GetSectionData(serverCfgSectionName, cfg)

	// 2.
	log.Println("Load server config data done")
}

func GetConfig() *Config {
	return cfg
}

func NewServer(h http.Handler) *http.Server {
	return &http.Server{
		Addr:           cfg.Host + ":" + cfg.Port,
		Handler:        h,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
