package models

import (
	"fmt"
	"gin-demo/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err error
		drive, dbName, user, password, host, tablePrefix string
	)
	// setting 包定义了Cfg 全局变量
	section ,err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get config section 'database' %v", err)
	}

	drive = section.Key("DRIVE").String()
	dbName = section.Key("NAME").String()
	user = section.Key("USER").String()
	password = section.Key("PASSWORD").String()
	host = section.Key("HOST").String()
	tablePrefix = section.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(drive, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			user,
			password,
			host,
			dbName,
		))
	if err !=nil {
		log.Println("connect error")
		log.Println(err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB()  {
	defer db.Close()
}