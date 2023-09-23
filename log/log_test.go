package log_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudtanj/Go-SwissKnife/env/value_types"
	"github.com/sudtanj/Go-SwissKnife/log"
	"testing"
)

type config struct {
}

func (c config) GetEnv() value_types.EnvConst {
	return value_types.Local
}

func TestLog_Initialize(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		inst := log.NewLog[config](config{})

		res := inst.Initialize()

		assert.NotNil(t, res)
		res.Info("zap successfully initialized")
	})
}
