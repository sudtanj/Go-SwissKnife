package viper_helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudtanj/Go-SwissKnife/env"
	"github.com/sudtanj/Go-SwissKnife/viper_helper"
	"os"
	"testing"
)

func Test_Initialize(t *testing.T) {
	t.Run("success load env", func(t *testing.T) {
		s, err := os.Getwd()
		assert.NoError(t, err)
		c := viper_helper.Initialize[TestCompositionConfig](s, "config.test")

		assert.Equal(t, c, TestCompositionConfig{
			RequiredConfig: env.RequiredConfig{Env: "local", Port: "8081", Addr: ""},
		})
	})
}
