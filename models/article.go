package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag Tag `json:"tag"`

	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	CreatedBy string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State int8 `json:"state"`
}

func (tag *Article) BeforeCreate(scope *gorm.Scope) error  {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

func (tag *Article) BeforeUpdate(scope *gorm.Scope) error  {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

func GetArticleById (id int) *Article {
	article := new(Article)
	db.First(article, id)
	return article
}

func ExistArticleById(id int) bool {
	article := GetArticleById(id)
	if article.ID > 0 {
		return true
	} else {
		return false
	}
}

func GetArticleCount (where map[string]interface{}) (count int) {
	db.Model(&Article{}).Where(where).Count(&count)
	return
}

func GetArticleList (page int, pageSize int, where map[string]interface{}) (article []Article) {
	db.Preload("Tag").Where(where).Offset(page).Limit(pageSize).Find(&article)
	return
}

func AddArticle (insert map[string]interface{}) bool {
	db.Create(&Article {
		TagID : insert["tag_id"].(int),
		Title : insert["title"].(string),
		Desc : insert["desc"].(string),
		Content : insert["content"].(string),
		CreatedBy : insert["created_by"].(string),
		State : insert["state"].(int8),
	})
	return true
}

func EditArticle (id int, update map[string]interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(update)
	return true
}

func DeleteArticle (id int) {
	db.Where("id = ?", id).Delete(Article{})
}

func CleanAllArticle () bool{
	db.Unscoped().Where("stata = ?", 0).Delete(&Article{})
	return true
}