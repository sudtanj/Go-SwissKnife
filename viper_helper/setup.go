package viper_helper

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/sudtanj/Go-SwissKnife/env"
)

func Initialize[T env.RequiredConfig](path string, fileName string) (config T) {
	viper.AddConfigPath(path)
	viper.SetConfigName(fileName)
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	conf := &T{}
	err = viper.Unmarshal(conf, func(config *mapstructure.DecoderConfig) {
		config.Squash = true
	})
	if err != nil {
		panic(fmt.Errorf("unable to decode into config struct, %v", err))
	}

	return *conf
}
