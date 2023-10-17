package main

import (
	"fmt"
	"gin-template/pkg/configx"
	"gin-template/pkg/logx"
	"gin-template/pkg/logx/zapx"
	"gin-template/pkg/mysqlx"
	"gin-template/pkg/redisx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
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

	config, err := initConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	mysqlConfig, _ := config.GetMysqlConfig()
	log.Println("mysql config", mysqlConfig)

	err = initMysql(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = initRedis(config)
	if err != nil {
		log.Fatalln(err.Error())
	}

	_ = initLog(config)
	logx.Info("haha")
	logx.Debug("debug")
	logx.Error("error")

	defer logx.Sync()

	//logging.Setup()
	//
	//// http shutdown
	//router := routers.InitRouter()
	//server := &http.Server{
	//	Addr:           fmt.Sprintf(":%d", conf.ServerConfig.HttpPort),
	//	Handler:        router,
	//	ReadTimeout:    conf.ServerConfig.ReadTimeout,
	//	WriteTimeout:   conf.ServerConfig.WriteTimeout,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//
	//go func() {
	//	if err := server.ListenAndServe(); err != nil {
	//		log.Printf("Listen: %s\n", err)
	//	}
	//}()
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//<-quit
	//
	//log.Println("Shutdown Server")
	//ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	//defer cancel()
	//if err := server.Shutdown(ctx); err != nil {
	//	log.Fatal("Server Shutdown:", err)
	//}

}

func initConfig() (*configx.ConfigX, error) {
	viperx, err := configx.NewViperX()
	if err != nil {
		return nil, fmt.Errorf("new config instance failed %w", err)
	}

	config := &configx.ConfigX{
		Driver: viperx,
	}

	return config, err
}

func initMysql(config *configx.ConfigX) error {
	mysqlConfig, err := config.GetMysqlConfig()
	if err != nil {
		return err
	}
	return mysqlx.New(mysqlConfig)
}

func initRedis(config *configx.ConfigX) error {
	redisConfig, err := config.GetRedisConfig()
	if err != nil {
		return err
	}
	return redisx.New(redisConfig)
}

func initLog(config *configx.ConfigX) error {

	logConfig, err := config.GetLogConfig()
	if err != nil {
		return err
	}

	encoder := zapx.GetZapCoreEncoder(logConfig.Encoder)
	tees := []zapx.TeeOption{
		{
			Ws: []zapcore.WriteSyncer{
				zapcore.AddSync(os.Stdout),
				zapcore.AddSync(zapx.GetHook(logConfig, logConfig.GetInfoPath())),
			},
			LevelEnablerFunc: func(level zapx.Level) bool {
				return level <= zap.InfoLevel
			},
			Encoder: encoder,
		},
		{
			Ws: []zapcore.WriteSyncer{
				zapcore.AddSync(os.Stdout),
				zapcore.AddSync(zapx.GetHook(logConfig, logConfig.GetErrPath())),
			},
			LevelEnablerFunc: func(level zapx.Level) bool {
				return level > zap.InfoLevel
			},
			Encoder: encoder,
		},
	}

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开发模式
	development := zap.Development()
	// 二次封装
	skip := zap.AddCallerSkip(1)

	logger := zapx.NewTee(tees, caller, development, skip)

	// defer logger.Sync()
	zapx.SetLogger(logger)
	return nil
}
