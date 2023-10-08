package rdb

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(
	isDebugMode bool,
	dbUser string,
	dbPass string,
	dbHost string,
	dbPort string,
	dbName string,
) error {
	if DB != nil {
		return nil
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbName)

	conf := &gorm.Config{PrepareStmt: true}
	if isDebugMode {
		conf.Logger = logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		)
	}

	conn, err := gorm.Open(mysql.Open(dsn), conf)
	if err != nil {
		return err
	}

	DB = conn
	return nil
}

func Close() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
