package env

import "github.com/sudtanj/Go-SwissKnife/env/value_types"

//go:generate mockgen -source=./config.go -destination=./mocks/config.go -package=mocks

type IConfig interface {
	GetEnv() value_types.EnvConst
	GetPort() string
	GetAddr() string
}

// do composition with this struct, this struct is design to work with viper
type RequiredConfig struct {
	Env  value_types.EnvConst `yaml:"env" validate:"required"`
	Port string               `yaml:"port" validate:"required"`
	Addr string               `yaml:"addr" validate:"required"`
}

func (c RequiredConfig) GetEnv() value_types.EnvConst {
	return c.Env
}

func (c RequiredConfig) GetPort() string {
	return c.Port
}

func (c RequiredConfig) GetAddr() string {
	return c.Addr
}
