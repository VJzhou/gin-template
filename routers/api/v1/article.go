package v1

import (
	"gin-demo/models"
	"gin-demo/pkg/app"
	"gin-demo/pkg/e"
	"gin-demo/pkg/logging"
	"gin-demo/pkg/setting"
	"gin-demo/pkg/util"
	"gin-demo/service/article_service"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticle (context *gin.Context ) {
	appg := app.Gin{context}
	id := com.StrTo(context.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appg.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{ID: id}
	article, err := articleService.GetArticle()
	if err != nil {
		appg.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	appg.Response(http.StatusOK, e.SUCCESS, article)
}

func GetArticles (context *gin.Context ) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	valid := validation.Validation{}
	appg := app.Gin{context}
	var state int = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state

		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	var tagId int = -1
	if arg := context.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appg.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	code := e.INVALID_PARAMS
	if ! valid.HasErrors() {
		code = e.SUCCESS

		data["lists"] = models.GetArticleList(util.GetPage(context), int(setting.AppConfig.PageSize), maps)
		data["total"] = models.GetArticleCount(maps)

	} else {
		for _, err := range valid.Errors {
			logging.Info("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}


	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func AddArticle (context *gin.Context ) {
	appg := app.Gin{context}
	tagId := com.StrTo(context.Query("tag_id")).MustInt()
	title := context.Query("title")
	desc := context.Query("desc")
	content := context.Query("content")
	createdBy := context.Query("created_by")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	image := context.Query("image")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appg.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	// 判断tag id 是否存在
	if ! models.ExistTagById(tagId) {
		appg.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}
	articleService := article_service.Article{
		TagID :  tagId,
		Title  :  title,
		Desc   :  desc,
		Content : content,
		CoverImageUrl : image,
		State :  state,
		CreatedBy : createdBy,
	}

	if !articleService.Add() {
		appg.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appg.Response(http.StatusOK, e.SUCCESS, nil)
}

func EditArticle (context *gin.Context ) {
	valid := validation.Validation{}

	id := com.StrTo(context.Param("id")).MustInt()
	tagId := com.StrTo(context.Query("tag_id")).MustInt()
	title := context.Query("title")
	desc := context.Query("desc")
	content := context.Query("content")
	modifiedBy := context.Query("modified_by")
	image := context.Query("image")

	var state int = -1
	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")

	appg := app.Gin{Context: context}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appg.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	if !models.ExistArticleById(id) {
		appg.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}
	// 判断tag id 是否存在
	if ! models.ExistTagById(tagId) {
		appg.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	article := article_service.Article{
		ID : id,
		TagID :  tagId,
		Title  :  title,
		Desc   :  desc,
		Content : content,
		CoverImageUrl : image,
		State :  state,
		ModifiedBy: modifiedBy,
	}
	if !article.Edit() {
		appg.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appg.Response(http.StatusOK, e.SUCCESS, nil)
	return
}

func DeleteArticle (context *gin.Context ) {
	id := com.StrTo(context.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	appg := app.Gin{Context: context}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appg.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	if !models.ExistArticleById(id) {
		appg.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article := article_service.Article{
		ID: id,
	}

	if ! article.Delete() {
		appg.Response(http.StatusOK, e.ERROR, nil)
		return
	}
	appg.Response(http.StatusOK, e.SUCCESS, nil)
	return
}


