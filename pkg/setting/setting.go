package setting

import (
	"log"
	"time"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	HTTPPort int

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	PageSize int8

	JwtSecret string
)

func init () {

	var err error
	Cfg , err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	runModeConfig()
	runServerConfig()
	runAppConfig()
}

func runModeConfig () {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func runServerConfig () {
	section , err:= Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	HTTPPort = section.Key("HTTP_PORT").MustInt(8080)
	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	ReadTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second

}

func runAppConfig () {
	section, err:= Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	JwtSecret = section.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = int8(section.Key("PAGE_SIZE").MustInt(10))
}

