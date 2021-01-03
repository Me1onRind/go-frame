package setting

import (
	"github.com/spf13/viper"
	"strings"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting(env string, dir string, configType string) (*Setting, error) {
	vp := viper.New()
	vp.SetConfigName(strings.ToLower(env))
	vp.AddConfigPath(dir)
	vp.SetConfigType(configType)
	if err := vp.ReadInConfig(); err != nil {
		return nil, err
	}
	return &Setting{vp}, nil
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	if err := s.vp.UnmarshalKey(k, v); err != nil {
		return err
	}
	return nil
}
