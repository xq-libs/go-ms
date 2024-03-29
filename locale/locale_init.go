package locale

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xq-libs/go-ms/config"
	"golang.org/x/text/language"
	"log"
)

const (
	cfgSectionName = "locale"
)

var (
	cfg    *Config
	bundle *i18n.Bundle
)

func init() {
	// 0.Do nothing without i18n config
	if !config.HasSection(cfgSectionName) {
		return
	}
	log.Printf("Will Load locale config data...")
	// 1.Load server config
	cfg = new(Config)
	config.GetSectionData(cfgSectionName, cfg)

	// 2.Load default language config
	defaultLang, err := language.Parse(cfg.Default)
	if err != nil {
		log.Println("Load locale default config language data failure")
		defaultLang = language.English
	}

	// 3.Create Bundle
	_bundle := i18n.NewBundle(defaultLang)
	_bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// 4.Load locale files
	for _, f := range cfg.Files {
		_bundle.MustLoadMessageFile(fmt.Sprintf("%s/%s", cfg.BaseDir, f))
	}
	bundle = _bundle
	log.Println("Load locale config data done.")
}

func GetConfig() *Config {
	return cfg
}

func NewLocalizer(languages ...string) *i18n.Localizer {
	return i18n.NewLocalizer(bundle, languages...)
}
