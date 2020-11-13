package v1

import (
	"gin-demo/models"
	"gin-demo/pkg/e"
	"gin-demo/pkg/setting"
	"gin-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

// @Summary get article tag list
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Success 200 {object} gin.H
// @Router /v1/tags/{id} [get]
func GetTag(context *gin.Context)  {
	name := context.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	var state int =  -1

	if arg := context.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS
	data["list"] = models.GetTagList(util.GetPage(context), int(setting.PageSize), maps)
	data["total"] = models.GetTagTotal(maps)

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

// @Summary add article tag
// @Produce  json
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} gin.H
// @Router /v1/tags/{id} [post]
func AddTag(context *gin.Context)  {
	name := context.Query("name")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	createdBy := context.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name,100, "name").Message("名称最长为100字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100,"created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors() {
		if models.ExistTagByName(name) {
			code = e.SUCCESS
			models.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
}

// @Summary Update article tag
// @Produce  json
// @Param id path int true "ID"
// @Param name body string true "Name"
// @Param state body int false "State"
// @Param modified_by body string true "ModifiedBy"
// @Success 200 {object} gin.H
// @Router /v1/tags/{id} [put]
func EditTag(context *gin.Context)  {
	id := com.StrTo(context.Query("id")).MustInt()
	name := context.Query("name")
	state := com.StrTo(context.DefaultQuery("state", "0")).MustInt()
	modifiedBy := context.Query("modified_by")

	// 校验数据
	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.MaxSize(name,100, "name").Message("名称最长为100字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100,"created_by").Message("修改人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	var code int = e.INVALID_PARAMS

	if !valid.HasErrors() {
		// 查询标签是否存在
		tag := models.GetTagById(id)
		if tag.ID <= 0 { // true 不存在标签
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			data := make(map[string]interface{})
			if name != "" {
				data["name"] = name
			}
			if modifiedBy != "" {
				data["modified_by"] = modifiedBy
			}
			data["state"] = state
			code = e.SUCCESS
			models.EditTag(tag, data)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}

func DeleteTag(context *gin.Context)  {
	id := com.StrTo(context.Query("id")).MustInt()

	tag := models.GetTagById(id)

	var code int
	if tag.ID > 0 {
		code = e.SUCCESS
		models.DeleteTag(tag)
	} else {
		code = e.ERROR_NOT_EXIST_TAG
	}
	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]interface{}),
	})
}
