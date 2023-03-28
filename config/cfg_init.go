package config

import (
	"github.com/xq-libs/go-ms/internal/os"
	"gopkg.in/ini.v1"
	"log"
)

const (
	defaultCfgFile = "conf/app.ini"
	fileEnvName    = "CONFIG_FILE"
)

var (
	cfgData *ini.File
)

func init() {
	// 1.Get config file
	cfgFile := os.GetEnvValue(fileEnvName, defaultCfgFile)
	log.Printf("Will load config data from file: %s", cfgFile)

	// 2.Load config data
	data, err := ini.Load(cfgFile)
	if err != nil {
		log.Panicf("Load config file data failure: %v", err)
	}
	cfgData = data
	log.Println("Load config file data done.")
}

func HasSection(name string) bool {
	return cfgData.HasSection(name)
}

func GetSectionData(name string, sectionData interface{}) {
	err := cfgData.Section(name).MapTo(sectionData)
	if err != nil {
		log.Panicf("Acquire config data with section: %s, failure: %v", name, err)
	}
}
