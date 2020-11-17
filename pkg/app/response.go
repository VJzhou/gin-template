package app

import (
	"gin-demo/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	Context *gin.Context
}

func (g *Gin) Response (httpcode , errCode int , data interface{}) {
	g.Context.JSON(httpcode, gin.H{
		"code" : errCode,
		"msg" : e.GetMsg(errCode),
		"data" : data,
	})
	return
}