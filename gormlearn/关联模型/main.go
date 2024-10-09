package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

	// 书籍表与出版社表一对多关系, 与作者表一对多

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

	//查询所有书籍
	var books []Book
	db.Model(&Book{}).Find(&books)
	fmt.Println(books)

	//查询(go语言大全)的出版社信息
	var book Book
	db.Model(&Book{}).Where("title = ?", "go语言大全").Preload("Publish").Find(&book)
	fmt.Println(book.Publish)

	//查询(go语言大全)的作者信息
	var book Book
	db.Model(&Book{}).Where("title = ?", "go语言大全").Preload("Authors").Find(&book)
	fmt.Println(book)

	//查询出版社信息和作者信息
	var book Book
	db.Model(&Book{}).Where("title = ?", "go语言大全").Preload("Authors").Preload("Publish").Find(&book)
	fmt.Println(book)

	db.Model(&Book{}).Where("title = ?", "go语言大全").Preload(clause.Associations).Find(&book)
	fmt.Println(book)

	//查询(西瓜出版社)写过的所有书
	var publish Publish
	db.Model(&Publish{}).Where("name = ?", "西瓜出版社").Preload("Books").Find(&publish)
	fmt.Println(publish.Books)

	//查询作者(杨阳)写过的所有书
	var authors Author
	db.Model(&Author{}).Where("name = ?", "杨阳").Preload("Books").Find(&authors)
	fmt.Println(authors)

	//自动关联, 一对多
	book := Book{
		Title:     "三国演义",
		Price:     1203,
		PublishID: 4,
		Publish: Publish{
			Name: "沈阳出版社",
		},
		Authors: nil,
	}

	db.Model(&Book{}).Create(&book)

	//自动关联多对多
	book := Book{
		Title:     "红楼梦",
		Price:     300,
		PublishID: 5,
		Authors:   []Author{{ID: 1}, {ID: 2}, {ID: 3}},
	}

	db.Model(&Book{}).Create(&book)
}
