package main

import (
	"fmt"
	"gin-template/config"
	"gin-template/pkg/logx"
	"gin-template/routers"
	"log"
	"net/http"
)

func main() {

	// TODO 热更新 创建子进程，将原进程退出

	config, err := config.New()
	if err != nil {
		log.Fatalln(err)
	}

	err = config.InitLog()
	if err != nil {
		log.Fatalln(err)
	}
	defer logx.Sync()

	err = config.InitMysql()
	if err != nil {
		logx.Error(err.Error())
	}

	err = config.InitRedis()
	if err != nil {
		logx.Error(err.Error())
	}

	svrConfig, err := config.GetServerConfig()
	if err != nil {
		logx.Error(err.Error())
	}

	// http shutdown
	router := routers.InitRouter()
	server := &http.Server{
		Addr:           fmt.Sprintf(":%s", svrConfig.HttpPort),
		Handler:        router,
		ReadTimeout:    svrConfig.ReadTimeout,
		WriteTimeout:   svrConfig.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	err = server.ListenAndServe()
	if err != nil {
		logx.Error(err.Error())
	}
}
