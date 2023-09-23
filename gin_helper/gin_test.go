package gin_helper_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudtanj/Go-SwissKnife/env"
	"github.com/sudtanj/Go-SwissKnife/gin_helper"
	"testing"
)

func TestGinHelper_GetAddr(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		inst := gin_helper.NewGinHelper[env.RequiredConfig](env.RequiredConfig{
			Port: "8080",
			Addr: "test.com",
		})

		assert.Equal(t, inst.GetAddr(), "test.com:8080")
	})
}

func TestGinHelper_GetEngine(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		inst := gin_helper.NewGinHelper[env.RequiredConfig](env.RequiredConfig{
			Port: "8080",
		})

		assert.NotNil(t, inst.GetEngine())
	})
}
