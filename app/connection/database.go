package connection

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DB_MAIN = "main"
const DB_TEST = "test"

type DBType string

func OpenDBConnection(dbType DBType, config *viper.Viper) (*gorm.DB, error) {
	var (
		gormLogger logger.Interface
		dialect    gorm.Dialector
	)
	switch dbType {
	case DB_TEST:
		{
			dialect = postgres.Open(config.GetString("database.test.dbUrl"))
		}
	case DB_MAIN:
		{
			dialect = postgres.Open(config.GetString("database.main.dbUrl"))
		}
	default:
		{
			return nil, errorhelper.ErrorCustom(500, "Invalid database type")
		}
	}

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
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(config.GetInt("database.dbMaxOpenConns"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.dbMaxIdleConns"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.dbConnMaxLifetime")) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.GetInt("database.dbConnMaxIdleTime")) * time.Minute)

	return db, nil
}
