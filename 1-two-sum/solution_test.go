package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	//assert.Equal(twoSum([]int{1}, 1))
	assert.Equal(t, [][2]int{[2]int{1, -1}}, twoSum([]int{1, -1}, 0))
	//twoSum([]int{1}, 1)
}
