package configx

import "github.com/go-ini/ini"

var _ Config = (*IniX)(nil)

type IniX struct {
	File *ini.File
}

func NewIniX(path string) (*IniX, error) {
	iniConfig, err := ini.Load(path)
	if err != nil {
		return nil, err
	}

	return &IniX{
		File: iniConfig,
	}, nil
}

func (i *IniX) ReadSection(key string, o interface{}) error {
	return i.File.Section(key).MapTo(o)
}
