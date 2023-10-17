package routers

import (
	"gin-template/conf"
	"gin-template/pkg/app"
	"gin-template/pkg/upload"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"net/http"
)

func InitRouter() *gin.Engine {
	route := gin.New()

	route.StaticFS("upload/images", http.Dir(upload.GetImageFullPath()))
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//route.POST("/upload", api.UploadImage)

	route.Use(gin.Logger())

	route.Use(gin.Recovery())

	gin.SetMode(conf.ServerConfig.RunMode)

	// 注册路由
	v1Api := route.Group("/v1")

	v1Api.GET("/name", func(c *gin.Context) {
		c.JSON(http.StatusOK, app.Response{
			Code: 1,
			Msg:  "asdasdasd",
		})
	})
	//v1Api.Use(jwt.JWT())

	return route
}
