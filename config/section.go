package config

import "time"

type Server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Mysql struct {
	User         string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type Redis struct {
	Host                  string
	Password              string
	DB                    int
	DialTimeout           time.Duration
	ReadTimeout           time.Duration
	WriteTimeout          time.Duration
	ContextTimeoutEnabled bool
	MaxRetries            int
	PoolSize              int
	PoolTimeout           time.Duration
	ConnMaxIdleTime       time.Duration
}
