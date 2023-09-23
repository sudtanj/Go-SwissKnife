package env

import "github.com/sudtanj/Go-SwissKnife/env/value_types"

//go:generate mockgen -source=./config.go -destination=./mocks/config.go -package=mocks

type IConfig interface {
	GetEnv() value_types.EnvConst
}
