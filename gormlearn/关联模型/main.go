package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

// 书籍类型
type Book struct {
	ID uint `gorm:"primarykey"`

	Title       string `gorm:"type:varchar(100);unique;not null"`
	Price       int64
	PublishDate *time.Time

	// 一对多
	PublishID int
	Publish   Publish
	// 多对多
	Authors []Author `gorm:"many2many:book_authors;"`
}

// 出版社类型
type Publish struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(100);unique;not null"`

	Books []Book
}

// 作者类型
type Author struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(100);unique;not null"`
	Age  int

	Books []Book `gorm:"many2many:book_authors;"`
}

func main() {

	// 书籍表与出版社表一对多关系, 与作者表多对多

	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&Author{})
	db.AutoMigrate(&Publish{})

}
