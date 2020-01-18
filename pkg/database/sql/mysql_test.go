package sql

import (
	"fmt"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type user struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func TestNewMysql(t *testing.T) {
	db, err := NewMysql(&MysqlConfig{
		DSN:             "root:root@/welfare_sign?charset=utf8&parseTime=True&loc=Local",
		IsPrintLog:      true,
		IsSingularTable: true,
	})
	if err != nil {
		panic(err)
	}

	defer db.Close()

	var u user
	if err := db.Find(&u).Error; err != nil {
		panic(err)
	}
	fmt.Println(u)
}
