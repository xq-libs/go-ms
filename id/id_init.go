package id

import (
	"log"

	"github.com/xq-libs/go-utils/codes"
	"github.com/xq-libs/go-utils/snowflake"
)

const (
	cfgSectionName = "sn"
)

var (
	sf *snowflake.SnowFlake
)

func init() {
	nsf, err := snowflake.NewSnowFlake(0)
	if err != nil {
		log.Panicf("Init Id Generator Failure: %v", err)
	}
	sf = nsf
}

func GenerateId() int64 {
	return sf.Generate()
}

func GenerateId62Str() string {
	return codes.EncodeIntBase62(sf.Generate())
}

func GenerateId32Str() string {
	return codes.EncodeIntBase32(sf.Generate())
}

func GenerateId16Str() string {
	return codes.EncodeIntBase16(sf.Generate())
}

func GenerateIdStr() string {
	return codes.EncodeIntBase10(sf.Generate())
}
