package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	// assert.Equal(t, 1, search([]int{2, 3, 4, 5, 0, 1}, 3))
	// assert.Equal(t, 4, search([]int{2, 3, 4, 5, 0, 1}, 0))
	// assert.Equal(t, 0, search([]int{6, 7, 8, 1, 2, 3, 4, 5}, 6))
	assert.Equal(t, 0, search([]int{1, 3, 5}, 1))
}
