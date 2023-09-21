package logx1

import (
	"path"
	"strings"
)

type Config struct {
	FilePath   string
	Encoder    string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

func (c *Config) Setup() {}

func (c *Config) GetInfoPath() string {
	prefix, suffix := getFileSuffixPrefix(c.FilePath)
	return path.Join(prefix + ".info" + suffix)
}

func (c *Config) GetErrPath() string {
	prefix, suffix := getFileSuffixPrefix(c.FilePath)
	return path.Join(prefix + ".err" + suffix)
}

// getFileSuffixPrefix 文件路径切割
func getFileSuffixPrefix(fileName string) (prefix, suffix string) {
	paths, _ := path.Split(fileName)
	base := path.Base(fileName)
	suffix = path.Ext(fileName)
	prefix = strings.TrimSuffix(base, suffix)
	prefix = path.Join(paths, prefix)
	return
}
