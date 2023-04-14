package registry

import (
	"github.com/xq-libs/go-ms/config"
	"log"
)

const (
	cfgClientSectionName   = "registry.client"
	cfgInstanceSectionName = "registry.instance"
)

var (
	cfg *Config
)

func init() {
	// 0.Do nothing without registry config
	if !(config.HasSection(cfgClientSectionName) && config.HasSection(cfgInstanceSectionName)) {
		return
	}
	// 1.Acquire registry config data
	log.Println("Load registry config data")
	cfg = new(Config)
	cfg.Client = &ClientConfig{}
	config.GetSectionData(cfgClientSectionName, cfg.Client)
	cfg.Instance = &InstanceConfig{}
	config.GetSectionData(cfgInstanceSectionName, cfg.Instance)

	// 3.Registry config done
	log.Println("Init registry done")
}

func GetConfig() *Config {
	return cfg
}
