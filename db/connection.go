package db

import (
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func OpenConnection(config *viper.Viper) *gorm.DB {
	var gormLogger logger.Interface
	dialect := postgres.Open(config.GetString("database.dbUrl"))

	switch config.GetString("database.logLevel") {
	case "error":
		gormLogger = logger.Default.LogMode(logger.Error)
	case "warn":
		gormLogger = logger.Default.LogMode(logger.Warn)
	case "info":
		gormLogger = logger.Default.LogMode(logger.Info)
	default:
		gormLogger = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: gormLogger,
	})
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxOpenConns(config.GetInt("database.dbMaxOpenConns"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.dbMaxIdleConns"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.dbConnMaxLifetime")) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.GetInt("database.dbConnMaxIdleTime")) * time.Minute)

	return db
}
