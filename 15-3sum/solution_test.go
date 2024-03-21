package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHard(t *testing.T) {
	assert.Equal(t, [][2]int{{1, -1}}, twoSum([]int{1, -1}, 0))
	assert.Equal(t, [][2]int{{1, -1}, {0, 0}}, twoSum([]int{1, -1, 0, 0}, 0))
	assert.Equal(t, [][2]int{{1, 1}, {-1, 3}}, twoSum([]int{1, 1, -1, 3}, 2))
}

func TestEasy(t *testing.T) {
	assert.Equal(t, []int{1, 2}, twoSumEasy([]int{3, 2, 4}, 6))
}
