package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPositive(t *testing.T) {
	assert.Equal(t, 8/2, divide(8, 2))
	assert.Equal(t, 10/3, divide(10, 3))
}

func TestNegative(t *testing.T) {
	assert.Equal(t, 7/-3, divide(7, -3))
	assert.Equal(t, -100/-3, divide(-100, -3))
	assert.Equal(t, -100/3, divide(-100, 3))
}

func TestLimits(t *testing.T) {
	// clamped
	assert.Equal(t, 2147483647, divide(-2147483648, -1))
}
