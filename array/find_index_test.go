package array_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/sudtanj/Go-SwissKnife/array"
	"testing"
)

func TestFindIndex_FindIndex(t *testing.T) {
	t.Run("return success", func(t *testing.T) {
		inst := array.NewFindIndex[string]()

		res := inst.FindIndex([]string{"test", "a", "b"}, func(comp string) bool {
			return comp == "a"
		})

		assert.Equal(t, res, 1)

		res1 := inst.FindIndex([]string{"test", "a", "b"}, func(comp string) bool {
			return comp == "abcd"
		})

		assert.Equal(t, res1, -1)
	})
}
