package viper_helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudtanj/Go-SwissKnife/viper_helper"
	"os"
	"testing"
)

func TestValidator_ValidateRequiredConfig(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		s, err := os.Getwd()
		assert.NoError(t, err)
		c := viper_helper.Initialize[TestCompositionConfig](s, "config.test")

		assert.PanicsWithError(t, "Key: 'TestCompositionConfig.RequiredConfig.Addr' Error:Field validation for 'Addr' failed on the 'required' tag", func() {
			viper_helper.ValidateRequiredConfig(c)
		})
	})
}
