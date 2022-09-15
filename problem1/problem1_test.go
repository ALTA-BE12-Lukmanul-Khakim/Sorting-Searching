package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMaxAndMin(t *testing.T) {
	t.Run("case 1", func(t *testing.T) {
		assert.Equal(t, "max: 8  index: 5 min: -2  index: 3", FindMaxAndMin([]int{5, 7, 4, -2, -1, 8}))
	})
	t.Run("case 2", func(t *testing.T) {
		assert.Equal(t, "max: 22  index: 3 min: -5  index: 1", FindMaxAndMin([]int{2, -5, -4, 22, 7, 7}))
	})

}
