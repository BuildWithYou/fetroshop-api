package connection

import (
	"time"

	"github.com/BuildWithYou/fetroshop-api/app/helper/errorhelper"
	loggerHelper "github.com/BuildWithYou/fetroshop-api/app/helper/logger"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const DB_MAIN = "main"
const DB_TEST = "test"

type DBType string

type Connection struct {
	DB  *gorm.DB
	Err error
}

func OpenDBConnection(dbType DBType, config *viper.Viper, lg *loggerHelper.Logger) *Connection {
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
			lg.Error("Invalid database type")
			return &Connection{
				Err: errorhelper.ErrorCustom(500, "Invalid database type"),
			}
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
		Logger:         gormLogger,
		TranslateError: true,
	})
	if err != nil {
		lg.UseError(err)
		return &Connection{
			Err: err,
		}
	}

	sqlDB, err := db.DB()
	if err != nil {
		lg.UseError(err)
		return &Connection{
			Err: err,
		}
	}

	sqlDB.SetMaxOpenConns(config.GetInt("database.dbMaxOpenConns"))
	sqlDB.SetMaxIdleConns(config.GetInt("database.dbMaxIdleConns"))
	sqlDB.SetConnMaxLifetime(time.Duration(config.GetInt("database.dbConnMaxLifetime")) * time.Minute)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.GetInt("database.dbConnMaxIdleTime")) * time.Minute)

	return &Connection{
		DB: db,
	}
}
