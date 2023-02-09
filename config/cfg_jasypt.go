package config

import (
	"github.com/xq-libs/go-ms/util/os"
	"github.com/xq-libs/go-utils/jasypt"
	"log"
)

const (
	jasyptPwdKey = "JASYPT_PASSWORD"
)

var (
	pwd string
	jas jasypt.Jasypt
)

func init() {
	pwd = os.GetEnvValue(jasyptPwdKey, "")
	if len(pwd) > 0 {
		j := jasypt.GetJasypt(jasypt.Options{
			Password: pwd,
		})
		jas = j
	}
}

func GetDecryptSectionData(name string, sectionData interface{}) {
	// 1.Load data obj
	GetSectionData(name, sectionData)

	// 2.Decrypt data obj
	if jas != nil {
		err := jasypt.DecryptObj(jas, sectionData)
		if err != nil {
			log.Panicf("Decrypt Load config data [%s] failure: %v", name, err)
		}
	}
}
