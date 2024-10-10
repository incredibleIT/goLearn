package main

import (
	"fmt"
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

	// find  找到书ID为3的所有作者
	var authors []Author
	db.Model(&Book{ID: 7}).Association("Authors").Find(&authors)
	fmt.Println(authors)

	// upsert 插入或更新
	db.Model(&Book{ID: 7}).Association("Authors").Append([]Author{{ID: 9, Name: "许仙", Age: 100}, {ID: 10, Name: "法海", Age: 1000}})

	//del 删除, clear()会将书籍ID为7的全部删除
	db.Model(&Book{ID: 7}).Association("Authors").Delete([]Author{{ID: 1}, {ID: 2}, {ID: 3}})

	// replace 更新, 会先执行Clear() 然后Append()
	db.Model(&Book{ID: 7}).Association("Authors").Replace([]Author{{ID: 1}, {ID: 2}, {ID: 3}})

	// Count计数, ID为7的书有几个作者
	db.Model(&Book{ID: 7}).Association("Authors").Count()

}
