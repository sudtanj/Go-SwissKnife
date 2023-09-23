package log

import (
	"github.com/sudtanj/Go-SwissKnife/conditional"
	"github.com/sudtanj/Go-SwissKnife/conditional/interfaces"
	"github.com/sudtanj/Go-SwissKnife/env"
	config_value_types "github.com/sudtanj/Go-SwissKnife/env/value_types"
	"go.uber.org/zap"
	zapAdapter "logur.dev/adapter/zap"
)

//go:generate ifacemaker -f ./log.go -s Log[T] -i "ILog[T any]" -p interfaces -o ./interfaces/log.go
//go:generate mockgen -source=./interfaces/log.go -destination=./mocks/log.go -package=mocks

type Log[T env.IConfig] struct {
	config            T
	loggerConditional interfaces.IConditional[*zap.Logger]
}

func NewLog[T env.IConfig](config T) *Log[T] {
	return &Log[T]{
		config,
		conditional.NewConditional[*zap.Logger](),
	}
}

func (l *Log[T]) Initialize() *zapAdapter.Logger {
	prodLog, err := zap.NewProduction()
	if err != nil {
		panic("error initialize zap prod")
	}
	devLog, err := zap.NewDevelopment()
	if err != nil {
		panic("error initialize zap dev")
	}
	zapInst := l.loggerConditional.If(
		l.config.GetEnv() == config_value_types.Prod,
		prodLog,
		devLog,
	)
	return zapAdapter.New(zapInst)
}
