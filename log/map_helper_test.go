package log

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestToMapInterface(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		err := errors.New("test")
		e := zap.Error(err)

		assert.Equal(t, ToMapInterface(e), LoggerMap{
			"error": err,
		})
	})
}
