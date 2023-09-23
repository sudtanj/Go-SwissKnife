package conditional_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudtanj/Go-SwissKnife/conditional"
	"testing"
)

func TestConditional_If(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		inst := conditional.NewConditional[string]()

		res := inst.If("t" == "a", "a", "b")

		assert.Equal(t, res, "b")

		res = inst.If("a" == "a", "a", "b")

		assert.Equal(t, res, "a")
	})
}
