package mysqlx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type MysqlConfig struct {
	User         string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

func (c *MysqlConfig) Dsn() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local", c.User, c.Password, c.Host, c.DBName, c.Charset, c.ParseTime)
	return dsn
}

func New(config *MysqlConfig) error {
	db, err := gorm.Open(mysql.Open(config.Dsn()), &gorm.Config{})
	if err != nil {
		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	//sqlDB.SetConnMaxLifetime(time.Hour)
	return nil
}

func GetDB() *gorm.DB {
	return db
}
