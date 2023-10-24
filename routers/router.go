package routers

import (
	"gin-template/config"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
	"net/http"
)

func InitRouter() *gin.Engine {
	route := gin.New()

	//route.StaticFS("upload/images", http.Dir(upload.GetImageFullPath()))
	//route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//route.POST("/upload", api.UploadImage)

	route.Use(gin.Logger())

	route.Use(gin.Recovery())

	gin.SetMode(config.ServerConfig().RunMode)

	// 注册路由
	v1Api := route.Group("/v1")

	v1Api.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//v1Api.Use(jwt.JWT())

	return route
}
