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
	localeCfg    *LocaleConfig
	localeBundle *i18n.Bundle
)

func init() {
	log.Printf("Load locale config data...")
	// 1.Load app config
	localeCfg = new(LocaleConfig)
	config.GetDecryptSectionData(cfgSectionName, localeCfg)

	// 2.Load default language config
	defaultLan, err := language.Parse(localeCfg.Default)
	if err != nil {
		log.Println("Load locale default config language data failure")
		defaultLan = language.English
	}
	i18n.NewBundle(defaultLan)

	// 3.Create Bundle
	bundle := i18n.NewBundle(defaultLan)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// 4.Load locale files
	for _, f := range localeCfg.Files {
		bundle.MustLoadMessageFile(fmt.Sprintf("%s/%s", localeCfg.BaseDir, f))
	}
	localeBundle = bundle
}

func NewLocalizer(languages ...string) *i18n.Localizer {
	return i18n.NewLocalizer(localeBundle, languages...)
}
