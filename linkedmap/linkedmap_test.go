package linkedmap

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedMap_PutAndGet(t *testing.T) {
	lm := NewLinkedMap[string, int](3)

	// Test inserting and retrieving items
	lm.Put("a", 1)
	lm.Put("b", 2)
	lm.Put("c", 3)

	val, ok := lm.Get("a")
	assert.True(t, ok, "expected to find key 'a'")
	assert.Equal(t, 1, val, "expected 'a' to be 1")

	val, ok = lm.Get("b")
	assert.True(t, ok, "expected to find key 'b'")
	assert.Equal(t, 2, val, "expected 'b' to be 2")

	val, ok = lm.Get("c")
	assert.True(t, ok, "expected to find key 'c'")
	assert.Equal(t, 3, val, "expected 'c' to be 3")
}

func TestLinkedMap_Remove(t *testing.T) {
	lm := NewLinkedMap[string, int](3)

	lm.Put("a", 1)
	lm.Put("b", 2)
	lm.Put("c", 3)

	lm.Remove("b")

	_, ok := lm.Get("b")
	assert.False(t, ok, "expected 'b' to be removed")

	val, ok := lm.Get("a")
	assert.True(t, ok, "expected to find key 'a'")
	assert.Equal(t, 1, val, "expected 'a' to be 1")

	val, ok = lm.Get("c")
	assert.True(t, ok, "expected to find key 'c'")
	assert.Equal(t, 3, val, "expected 'c' to be 3")
}

func TestLinkedMap_UpdateValue(t *testing.T) {
	lm := NewLinkedMap[string, int](3)

	lm.Put("a", 1)
	lm.Put("a", 10) // Update the value for key "a"

	// Check if the value is updated
	val, ok := lm.Get("a")
	assert.True(t, ok, "expected to find key 'a'")
	assert.Equal(t, 10, val, "expected 'a' to be 10")
}

func TestGetMostRecent(t *testing.T) {
	lm := NewLinkedMap[string, int](3)
	lm.Put("a", 1)
	lm.Put("b", 2)

	val, ok := lm.GetNewest()
	assert.Equal(t, "b", val.key)
	assert.Equal(t, 2, val.value)
	assert.True(t, ok)

	_, _ = lm.Get("a")
	val, ok = lm.GetNewest()
	assert.Equal(t, "a", val.key)
	assert.Equal(t, 1, val.value)
	assert.True(t, ok)
}

func TestGetLeastRecent(t *testing.T) {
	lm := NewLinkedMap[string, int](3)
	lm.Put("a", 1)
	lm.Put("b", 2)

	val, ok := lm.GetOldest()
	assert.Equal(t, "a", val.key)
	assert.Equal(t, 1, val.value)
	assert.True(t, ok)

	_, _ = lm.Get("a")
	val, ok = lm.GetOldest()
	assert.Equal(t, "b", val.key)
	assert.Equal(t, 2, val.value)
	assert.True(t, ok)
}
