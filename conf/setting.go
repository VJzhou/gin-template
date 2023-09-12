package conf

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	JwtSecret       string
	PageSize        int
	RuntimeRootPath string

	ImagePrefixPath string
	ImageSavePath   string
	ImagaMaxSize    int
	ImageAllowExts  []string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppConfig = &App{}

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = &Server{}

type DataBase struct {
	Drive       string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var DatabaseConfig = &DataBase{}

type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

var RedisConfig1 = &Redis{}

func Setup() {
	config, err := ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	err = config.Section("app").MapTo(AppConfig)
	AppConfig.ImagaMaxSize = AppConfig.ImagaMaxSize * 1024 * 1024

	err = config.Section("server").MapTo(ServerConfig)
	if err != nil {
		log.Fatalf("config.mapTo serverconfig err: %v", err)
	}
	ServerConfig.ReadTimeout = ServerConfig.ReadTimeout * time.Second
	ServerConfig.WriteTimeout = ServerConfig.WriteTimeout * time.Second

	err = config.Section("database").MapTo(DatabaseConfig)
	if err != nil {
		log.Fatalf("config.mapto DatabaseSetting err: %v", err)
	}

	err = config.Section("redis").MapTo(RedisConfig1)
	if err != nil {
		log.Fatalf("config.mapto RedisSetting err: %v", err)
	}
}

//
//
//var (
//	Cfg *ini.File
//
//	RunMode string
//
//	HTTPPort int
//
//	ReadTimeout time.Duration
//
//	WriteTimeout time.Duration
//
//	PageSize int8
//
//	JwtSecret string
//)
//
//func init () {
//
//	var err error
//	Cfg , err = ini.Load("conf/app.ini")
//	if err != nil {
//		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
//	}
//	runModeConfig()
//	runServerConfig()
//	runAppConfig()
//}
//
//func runModeConfig () {
//	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
//}
//
//func runServerConfig () {
//	section , err:= Cfg.GetSection("server")
//	if err != nil {
//		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
//	}
//
//	HTTPPort = section.Key("HTTP_PORT").MustInt(8080)
//	ReadTimeout = time.Duration(section.Key("READ_TIMEOUT").MustInt(60)) * time.Second
//	ReadTimeout = time.Duration(section.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
//
//}
//
//func runAppConfig () {
//	section, err:= Cfg.GetSection("app")
//	if err != nil {
//		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
//	}
//
//	JwtSecret = section.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
//	PageSize = int8(section.Key("PAGE_SIZE").MustInt(10))
//}
