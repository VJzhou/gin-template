package routers

import (
	"gin-demo/middleware/jwt"
	"gin-demo/pkg/setting"
	"gin-demo/pkg/upload"
	"gin-demo/routers/api"
	v1 "gin-demo/routers/api/v1"
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
	route.POST("/upload", api.UploadImage)

	route.Use(gin.Logger())

	route.Use(gin.Recovery())

	gin.SetMode(setting.ServerConfig.RunMode)

	//route.GET("/test", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message" : "test",
	//	})
	//})
	route.GET("/auth", api.GetToken)

	// 注册路由
	v1Api := route.Group("/v1")
	v1Api.Use(jwt.JWT())
	{
		v1Api.GET("/tag", v1.GetTag)
		v1Api.POST("/tag", v1.AddTag)
		v1Api.PUT("/tag/:id", v1.EditTag)
		v1Api.DELETE("/tag/:id", v1.DeleteTag)

		v1Api.GET("/article/:id", v1.GetArticle)
		v1Api.GET("/articles", v1.GetArticles)
		v1Api.POST("/article", v1.AddArticle)
		v1Api.PUT("/article/:id", v1.EditArticle)
		v1Api.DELETE("/article/:id", v1.DeleteArticle)
	}

	return route
}
