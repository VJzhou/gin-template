package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model // 内嵌类型

	Name string
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int8 `json:"state"`
}

func GetTagList(page int , pageSize int, where interface{}) (tags []Tag) {
	db.Where(where).Offset(page).Limit(pageSize).Find(&tags)
	return
}

func GetTagTotal(where interface{}) (count int) {
	db.Model(&Tag{}).Where(where).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).Find(&tag)

	if tag.ID > 0 {
		return false
	}
	return true
}

func AddTag (name string, state int, createdBy string) bool{
	db.Create(&Tag{
		Name: name,
		State: int8(state),
		CreatedBy: createdBy,
	})
	return true
}

func EditTag (tag *Tag, update map[string]interface{}) bool {
	db.Model(tag).Update(update)
	return true
}

func GetTagById (id int) *Tag {
	tag := new(Tag)
	tag.ID = id
	db.First(tag)
	return tag
}

func ExistTagById(id int) bool {
	tag := GetTagById(id)
	if tag.ID > 0 {
		return true
	} else {
		return false
	}
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}


func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func DeleteTag (tag *Tag) {
	db.Delete(tag)
}
