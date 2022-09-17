package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximumBuyProduct(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		assert.Equal(t, 3, MaximumBuyProduct(50000, []int{25000, 25000, 10000, 14000}))
	})
	t.Run("case 2", func(t *testing.T) {
		assert.Equal(t, 4, MaximumBuyProduct(30000, []int{15000, 10000, 12000, 5000, 3000}))
	})

}
