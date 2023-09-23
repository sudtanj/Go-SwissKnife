package env

import "github.com/sudtanj/Go-SwissKnife/env/value_types"

//go:generate mockgen -source=./config.go -destination=./mocks/config.go -package=mocks

type IConfig interface {
	GetEnv() value_types.EnvConst
}

// do composition with this struct, this struct is design to work with viper
type RequiredConfig struct {
	Env string `yaml:"env"`
}

func (c *RequiredConfig) GetEnv() string {
	return c.Env
}
