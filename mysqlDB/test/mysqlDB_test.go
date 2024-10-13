package DBtest

import (
	creatConnection "Toconnect/utils"
	"fmt"
	"gorm.io/gorm"
	"testing"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func TestCreateMysqlConnection(t *testing.T) {

	dbViper := creatConnection.GetConfig("mysql")

	conn := creatConnection.CreateMysqlConnection(dbViper.GetString("db.username"),
		dbViper.GetString("db.password"),
		dbViper.GetString("db.host"),
		dbViper.GetInt("db.port"))

	if err := conn.AutoMigrate(&User{}); err != nil {
		panic(fmt.Errorf("表创建失败: %s", err))
	}

}
