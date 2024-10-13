package creatConnection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var myLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	logger.Config{
		SlowThreshold:             time.Second, // Slow SQL threshold
		LogLevel:                  logger.Info, // Log level
		IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
		ParameterizedQueries:      true,        // Don't include params in the SQL log
		Colorful:                  true,        // Disable color
	},
)

func CreateMysqlConnection(dbname string, password string, host string, port int) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/gorm?charset=utf8mb4&parseTime=True&loc=Local", dbname, password, host, port)

	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:      myLogger,
		PrepareStmt: true,
	})

	if err != nil {
		panic("failed to connect database")
	}

	sqlDB, _ := db.DB()

	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetMaxIdleConns(20)

	return db

}
