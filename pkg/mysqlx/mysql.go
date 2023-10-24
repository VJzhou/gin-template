package mysqlx

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

type Config struct {
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

func (dsn *Config) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dsn.User, dsn.Password, dsn.Host, dsn.DBName, dsn.Charset, dsn.ParseTime)
}

func New(config *Config) error {
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
