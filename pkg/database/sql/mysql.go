package sql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pkg/errors"
)

// MysqlConfig mysql config
type MysqlConfig struct {
	DSN             string // db conn string
	IsPrintLog      bool   // enable log mode
	IsSingularTable bool   // enable singular table
}

// NewMysql new mysql db conn
func NewMysql(conf *MysqlConfig) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", conf.DSN)
	if err != nil {
		return nil, errors.WithMessage(err, "mysql opne error")
	}
	if err := db.DB().Ping(); err != nil {
		return nil, errors.WithMessage(err, "mysql ping error")
	}
	db.LogMode(conf.IsPrintLog)
	db.SingularTable(conf.IsSingularTable)

	return db, nil
}
