package configx

import (
	"github.com/spf13/viper"
)

var _ Config = (*ViperX)(nil)

type ViperX struct {
	vp *viper.Viper
}

// TODO Modified
func NewViperX() (*ViperX, error) {
	vp := viper.New()
	vp.SetConfigName("config") // 设置配置文件名称
	vp.AddConfigPath("etc/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	return &ViperX{vp}, nil
}

func (s *ViperX) ReadSection(key string, v interface{}) error {
	return s.vp.UnmarshalKey(key, v)
}
