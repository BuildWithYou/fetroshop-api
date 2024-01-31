package connection

import (
	"fmt"
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
			dbName := config.GetString("database.test.dbName")
			dbUsername := config.GetString("database.test.dbUsername")
			dbPassword := config.GetString("database.test.dbPassword")
			dbHost := config.GetString("database.test.dbHost")
			dbPort := config.GetInt("database.test.dbPort")
			dbParam := config.GetString("database.test.dbParam")
			dbUrl := fmt.Sprintf(
				"postgresql://%s:%s@%s:%d/%s?%s",
				dbUsername,
				dbPassword,
				dbHost,
				dbPort,
				dbName,
				dbParam,
			)
			dialect = postgres.Open(dbUrl)
		}
	case DB_MAIN:
		{
			dbName := config.GetString("database.main.dbName")
			dbUsername := config.GetString("database.main.dbUsername")
			dbPassword := config.GetString("database.main.dbPassword")
			dbHost := config.GetString("database.main.dbHost")
			dbPort := config.GetInt("database.main.dbPort")
			dbParam := config.GetString("database.main.dbParam")
			dbUrl := fmt.Sprintf(
				"postgresql://%s:%s@%s:%d/%s?%s",
				dbUsername,
				dbPassword,
				dbHost,
				dbPort,
				dbName,
				dbParam,
			)
			dialect = postgres.Open(dbUrl)
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
