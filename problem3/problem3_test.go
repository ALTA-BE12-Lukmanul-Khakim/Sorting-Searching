package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayingDomino(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, []int{3, 4}, PlayingDomino([][]int{[]int{6, 5}, []int{3, 4}, []int{2, 1}, []int{3, 3}}, []int{4, 3}))
	})
	t.Run("case 2", func(t *testing.T) {
		assert.Equal(t, []int(nil), PlayingDomino([][]int{[]int{6, 6}, []int{2, 4}, []int{3, 6}}, []int{5, 1}))
	})

}
