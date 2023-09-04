package redisx

import (
	"encoding/json"
	"gin-demo/conf"
	"github.com/gomodule/redigo/redis"
	"time"
)

var Conn *redis.Pool

func Setup() error {
	Conn = &redis.Pool{
		MaxIdle:     conf.RedisConfig.MaxIdle,
		MaxActive:   conf.RedisConfig.MaxActive,
		IdleTimeout: conf.RedisConfig.IdleTimeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.RedisConfig.Host)
			if err != nil {
				return nil, err
			}
			if conf.RedisConfig.Password != "" {
				if _, err := c.Do("AUTH", conf.RedisConfig.Password); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return nil
}

func Set(key string, data interface{}, time int) error {
	conn := Conn.Get()
	defer conn.Close()

	value, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", key, value)
	if err != nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err != nil {
		return err
	}
	return nil
}

func Exists(key string) bool {
	conn := Conn.Get()
	defer conn.Close()

	exist, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false
	}
	return exist
}

func Get(key string) ([]byte, error) {
	conn := Conn.Get()
	defer conn.Close()

	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func Delete(key string) (bool, error) {
	conn := Conn.Get()
	defer conn.Close()

	return redis.Bool(conn.Do("DEL", key))
}

func LikeDeletes(key string) error {
	conn := Conn.Get()
	defer conn.Close()

	keys, err := redis.Strings(conn.Do("KEYS", "*"+key+"*"))
	if err != nil {
		return err
	}

	for _, key := range keys {
		_, err = Delete(key)
		if err != nil {
			return err
		}
	}

	return nil
}
