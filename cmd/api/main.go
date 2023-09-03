package main

import (
	"context"
	"fmt"
	"gin-demo/models"
	"gin-demo/pkg/logging"
	"gin-demo/pkg/redisx"
	"gin-demo/pkg/setting"
	"gin-demo/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// @title Docker监控服务
// @version 1.0
// @description docker监控服务后端API接口文档

// @contact.name API Support
// @contact.url http://www.swagger.io/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host 127.0.0.1:8088
func main() {

	// 热更新 创建子进程，将原进程退出
	//endless.DefaultReadTimeOut = setting.ReadTimeout
	//endless.DefaultWriteTimeOut = setting.WriteTimeout
	//endless.DefaultMaxHeaderBytes = 1 << 20
	//endPoint := fmt.Sprintf(":%d", setting.HTTPPort)
	//
	//server := endless.NewServer(endPoint, routers.InitRouter())
	//server.BeforeBegin = func(add string) {
	//  log.Printf("Actual pid is %d", syscall.Getpid())
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//  log.Printf("Server err: %v", err)
	//}

	setting.Setup()
	models.Setup()
	logging.Setup()
	if err := redisx.Setup(); err != nil {
		logging.Error(err)
	}
	// http shutdown
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerConfig.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerConfig.ReadTimeout,
		WriteTimeout:   setting.ServerConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")
}
