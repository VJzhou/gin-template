package models

import (
	"fmt"
	"gin-demo/pkg/setting"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseConfig.Drive, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseConfig.User,
		setting.DatabaseConfig.Password,
		setting.DatabaseConfig.Host,
		setting.DatabaseConfig.Name,
		))
	if err !=nil {
		log.Println("connect error")
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseConfig.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func CloseDB()  {
	defer db.Close()
}

func updateTimeStampCreateCallback (scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
				createTimeField.Set(nowTime)
			}
		}
		if updateTimeField, ok := scope.FieldByName("UpdatedOn"); ok {
			if updateTimeField.IsBlank {
				updateTimeField.Set(nowTime)
			}
		}
	}
}

func updateTimeStampForUpdateCallback (scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}