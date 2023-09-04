package mysqlx

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type MysqlX struct{}

// TODO Modified
func New(dsn string) error {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func GetDB() *gorm.DB {
	return db
}
