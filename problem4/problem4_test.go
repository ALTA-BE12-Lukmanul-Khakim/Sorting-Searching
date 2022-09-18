package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMostAppearItems(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, "golang -> 1 ruby -> 2 js -> 4 ", MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"}))
	})
	t.Run("case 2", func(t *testing.T) {
		assert.Equal(t, "C -> 1 D -> 2 B -> 3 A -> 4 ", MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"}))
	})

}
