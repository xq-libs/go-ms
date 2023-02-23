package database

import (
	"fmt"
	"github.com/xq-libs/go-ms/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const (
	cfgSectionName = "db"
)

var (
	cfg *Config
	db  *gorm.DB
)

func init() {
	// 1.Acquire db config data
	cfg = new(Config)
	config.GetDecryptSectionData(cfgSectionName, cfg)

	// 2.Create db dsn
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)

	// 3.Create db connections
	_db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		log.Panicf("Connection to db failure: %v", err)
	}

	// 4.Config db env
	sqlDb, _ := _db.DB()
	sqlDb.SetMaxOpenConns(50)
	sqlDb.SetMaxIdleConns(5)

	db = _db
	log.Println("Init db done")
}

func GetConfig() *Config {
	return cfg
}

func GetDb() *gorm.DB {
	return db
}
