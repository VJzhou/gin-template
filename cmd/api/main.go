package main

import (
	"fmt"
	"gin-template/conf"
	"gin-template/pkg/configx"
	"gin-template/pkg/logx"
	"gin-template/pkg/logx/zapx"
	"gin-template/pkg/mysqlx"
	"gin-template/pkg/redisx"
	"gin-template/routers"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"os"
)

func main() {

	// TODO 热更新 创建子进程，将原进程退出

	config, err := initConfig()
	if err != nil {
		log.Fatalln(err.Error())
	}

	err = initLog(config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer logx.Sync()

	err = initMysql(config)
	if err != nil {
		logx.Error(err.Error())
	}

	err = initRedis(config)
	if err != nil {
		logx.Error(err.Error())
	}

	svrConfig, err := initServer(config)
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

func initServer(config *configx.ConfigX) (*conf.ServerConfigX, error) {
	serverConfig, err := config.GetServerConfig()
	if err != nil {
		return nil, err
	}
	return serverConfig, nil
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

	zapx.SetLogger(logger)
	return nil
}
