package creatConnection

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"syscall"
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

// 返回文件的创建时间，最后访问时间，最后修改时间。均为时间戳，精确到秒
func GetTimeInfo(stat os.FileInfo) (int64, int64, int64) {
	// windows下代码如下
	attr := stat.Sys().(*syscall.Win32FileAttributeData)
	return attr.CreationTime.Nanoseconds() / int64(1e9),
		attr.LastAccessTime.Nanoseconds() / int64(1e9),
		attr.LastWriteTime.Nanoseconds() / int64(1e9)
}

func RotateDBLog(db *gorm.DB, filePath string) error {
	if db == nil || len(filePath) == 0 {
		return nil
	}

	dbFile := RootPath + "log/" + filePath // 拿到dblog文件
	// 如果存在, 或者不存在
	fileInfo, err := os.Stat(dbFile)
	if err == nil { // 存在
		ct, _, _ := GetTimeInfo(fileInfo)
		creatTime := time.Unix(ct, 0)
		now := time.Now()

		if creatTime.Year() != now.Year() || creatTime.YearDay() != now.YearDay() {
			// 证明不是一天了
			// 首先重命名文件
			err = os.Rename(dbFile, dbFile+"."+creatTime.Format("2006-01-02 15:04:05"))

			if err != nil {
				return err
			} else {
				fmt.Println("重命名成功")
			}
		}
	} else {
		if os.IsNotExist(err) { // 不存在

		} else { // 其他错误
			return err
		}
	}

	// 不存在, 或者不是同一天
	newFile, err := os.OpenFile(dbFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	} else {
		db.Logger = logger.New(
			log.New(newFile, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      true,
			})
	}

	return nil
}

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
