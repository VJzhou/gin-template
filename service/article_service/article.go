package article_service

import (
	"encoding/json"
	"fmt"
	"gin-demo/models"
	"gin-demo/pkg/gredis"
	"gin-demo/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

func (a *Article) GetArticle () (*models.Article, error) {
	var modelArticle *models.Article
	// 读取缓存
	cacheArticle := cache_service.Article{ID: a.ID}
	key := cacheArticle.GetArticleKey()
	if gredis.Exists(key) {
		data , err := gredis.Get(key)
		if err != nil {
			return nil, err
		}
		json.Unmarshal(data, modelArticle)
		return modelArticle, nil
	}
	// 查询db
	fmt.Println("cache......")
	dbArticle := models.GetArticleById(a.ID)

	if dbArticle.ID <= 0 {
		return dbArticle, nil
	}

	gredis.Set(key, dbArticle, 3600)
	return dbArticle, nil
}

func (a *Article) Add () bool {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"image":           a.CoverImageUrl,
		"state":           int8(a.State),
	}
	if !models.AddArticle(article){
		return false
	}
	return true
}

func (a *Article) Edit () bool {
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"modified_by":      a.ModifiedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           int8(a.State),
	})
}

func  (a *Article) Delete() bool {
	models.DeleteArticle(a.ID)
	return true
}

func (a *Article) IsExist(article *models.Article) bool {
	if article.ID > 0 {
		return true
	}
	return false
}