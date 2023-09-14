package main

import (
	"fmt"
	"gin-demo/pkg/configx"
	"gin-demo/pkg/mysqlx"
	"gin-demo/pkg/redisx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
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

	conf := zap.NewProductionConfig()

	conf.Encoding = "console"

	format := "[%s]"

	conf.EncoderConfig.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf(format, t.Format("2006-01-02 15:04:05")))
	}

	conf.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf(format, caller.TrimmedPath()))
	}

	conf.EncoderConfig.EncodeLevel = func(level zapcore.Level, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(fmt.Sprintf(format, level.String()))
	}

	logger, _ := conf.Build()
	logger.Info("service start")

	logger.Info("info msg", zap.String("haha", "name"), zap.Int("age", 18), zap.Duration("timer", time.Minute))

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

	log.Println("Server exiting")
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
