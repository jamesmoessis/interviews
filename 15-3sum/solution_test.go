package main

import (
	"slices"
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

func Test3Sum(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func TestSortingArray(t *testing.T) {
	arr := [3]int{3, 2, 1}
	slices.Sort(arr[:])
	assert.Equal(t, [3]int{1, 2, 3}, arr)
}
