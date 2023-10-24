package config

import (
	"errors"
	"fmt"
	"gin-template/pkg/configx"

	"gin-template/pkg/logx/zapx"
	"gin-template/pkg/mysqlx"
	"gin-template/pkg/redisx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	server *Server
	mysql  *mysqlx.Config
	redis  *redisx.Config
	log    *zapx.Config
)

type ConfigX struct {
	Driver configx.Config
}

func New() (*ConfigX, error) {
	viperx, err := configx.NewViperX()
	if err != nil {
		return nil, fmt.Errorf("new config instance failed %w", err)
	}

	return &ConfigX{viperx}, nil
}

func (c *ConfigX) InitLog() error {
	logConfig, err := c.GetLogConfig()
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

func (c *ConfigX) InitMysql() error {
	mysqlConfig, err := c.GetMysqlConfig()
	if err != nil {
		return err
	}

	config, ok := mysqlConfig.(*mysqlx.Config)
	if !ok {
		return errors.New("can not convert to mysqlx config")
	}

	mysql = config

	return mysqlx.New(config)
}

func (c *ConfigX) InitRedis() error {
	redisConfig, err := c.GetRedisConfig()
	if err != nil {
		return err
	}

	config, ok := redisConfig.(*redisx.Config)
	if !ok {
		return errors.New("can not convert to redisx config")
	}

	redis = config

	return redisx.New(config)
}

func (c *ConfigX) GetServerConfig() (*Server, error) {
	var config Server
	err := c.Driver.ReadSection("Server", &config)
	if err != nil {
		return nil, err
	}

	server = &config

	return server, nil
}

func (c *ConfigX) GetMysqlConfig() (interface{}, error) {
	var mysqlConfig Mysql
	err := c.Driver.ReadSection("Mysql", &mysqlConfig)
	if err != nil {
		return nil, fmt.Errorf("get mysql config failed %w", err)
	}
	return &mysqlConfig, nil
}

func (c *ConfigX) GetRedisConfig() (interface{}, error) {
	var redisConfig Redis
	err := c.Driver.ReadSection("Redis", &redisConfig)
	if err != nil {
		return nil, err
	}
	return &redisConfig, nil
}

func (c *ConfigX) GetLogConfig() (*zapx.Config, error) {
	var config zapx.Config
	err := c.Driver.ReadSection("Log", &config)
	if err != nil {
		return nil, err
	}

	log = &config

	return log, nil
}

func RedisConfig() *redisx.Config {
	return redis
}

func MysqlConfig() *mysqlx.Config {
	return mysql
}

func LogConfig() *zapx.Config {
	return log
}

func ServerConfig() *Server {
	return server
}
