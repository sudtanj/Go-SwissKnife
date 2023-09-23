package env

import "github.com/sudtanj/Go-SwissKnife/env/value_types"

//go:generate mockgen -source=./config.go -destination=./mocks/config.go -package=mocks

type IConfig interface {
	GetEnv() value_types.EnvConst
}

// do composition with this struct, this struct is design to work with viper
type RequiredConfig struct {
	Env value_types.EnvConst `yaml:"env"`
}

func (c RequiredConfig) GetEnv() value_types.EnvConst {
	return c.Env
}