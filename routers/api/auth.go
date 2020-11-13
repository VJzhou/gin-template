package api

import (
	"gin-demo/models"
	"gin-demo/pkg/e"
	"gin-demo/pkg/logging"
	"gin-demo/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

// @Summary Get token
// @Produce json
// @Param username path string true "USERNAME"
// @Param password path string true "PASSWORD"
// @Success 200 {object} gin.H{}
// @Router /auth [get]
func GetToken (context *gin.Context) {
	username := context.Query("username")
	password := context.Query("password")

	valid := validation.Validation{}

	auth := auth{Username: username, Password: password}

	ok, _ := valid.Valid(&auth)

	data := make(map[string]interface{})
	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token , err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}
	} else {
		for _, err := range valid.Errors {
			logging.Info(err.Key, err.Message)
		}
	}

	context.JSON(http.StatusOK, gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}
