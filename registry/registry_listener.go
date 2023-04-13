package registry

import (
	"log"
	"net/http"
	"time"
)

type RegisterListener struct {
}

var defaultListener = RegisterListener{}

func GetDefaultListener() *RegisterListener {
	return &defaultListener
}

func (l *RegisterListener) OnStart(s *http.Server) {
	log.Println("Will register current service to registry center.")
	// 1. Register current instance.
	go registerService()
	// 2. Load all instance.
	go loadAllServices()
}

func (l *RegisterListener) OnShutdown() {
	log.Println("Will remove current service from registry center.")
	unregisterService()
}

func registerService() {
	clientConfig := GetConfig().Client
	for i := 0; i < clientConfig.TryTimes; i++ {
		log.Printf("Register current service on %d times.\n", i)
		result := RegisterServiceInstance()
		if result {
			log.Printf("Register success.")
			return
		}
		time.Sleep(time.Duration(clientConfig.RefreshInterval) * time.Second)
	}
}

func loadAllServices() {
	clientConfig := GetConfig().Client
	for {
		LoadAllServiceInstances()
		time.Sleep(time.Duration(clientConfig.RefreshInterval) * time.Second)
	}
}

func unregisterService() {
	if !UnregisterServiceInstance() {
		log.Printf("Unregister current service failure.")
	}
}
