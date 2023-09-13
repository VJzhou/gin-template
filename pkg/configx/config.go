package configx

import (
	"fmt"
	"gin-demo/pkg/mysqlx"
	"gin-demo/pkg/redisx"
)

type Config interface {
	ReadSection(string, interface{}) error
}

type ConfigX struct {
	Driver Config
}

func (c *ConfigX) GetMysqlConfig() (*mysqlx.MysqlConfig, error) {
	var mysqlConfig mysqlx.MysqlConfig
	err := c.Driver.ReadSection("Mysql", &mysqlConfig)
	if err != nil {
		return nil, fmt.Errorf("get mysql config failed %w", err)
	}
	return &mysqlConfig, nil
}

func (c *ConfigX) GetRedisConfig() (*redisx.RedisConfig, error) {
	var redisConfig redisx.RedisConfig
	err := c.Driver.ReadSection("Redis", &redisConfig)
	if err != nil {
		return nil, err
	}
	return &redisConfig, nil
}
