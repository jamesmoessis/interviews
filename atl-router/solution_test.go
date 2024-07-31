package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAndCallRoute(t *testing.T) {
	r := NewRouter()
	r.AddRoute("/foo/bar", "a")
	r.AddRoute("/foo/baz", "b")
	r.AddRoute("/abc/def", "c")
	r.AddRoute("/abc", "d")

	assert.Equal(t, r.CallRoute("/foo/bar"), Handler("a"))
	assert.Equal(t, r.CallRoute("/foo/baz"), Handler("b"))
	assert.Equal(t, r.CallRoute("/abc/def"), Handler("c"))
	assert.Equal(t, r.CallRoute("/abc"), Handler("d"))
	assert.Equal(t, r.CallRoute("/"), Handler("404 not found"))
	assert.Equal(t, r.CallRoute("/d"), Handler("404 not found"))
}

func TestWildcards(t *testing.T) {
	r := NewRouter()
	r.AddRoute("/foo/*", "a")
	r.AddRoute("/foo/*/baz", "b")
	r.AddRoute("/*", "c")
	r.AddRoute("/*/bar", "d")

	assert.Equal(t, r.CallRoute("/foo/abc"), Handler("a"))
	assert.Equal(t, r.CallRoute("/foo/abc/baz"), Handler("b"))
	assert.Equal(t, r.CallRoute("/hello"), Handler("c"))
	assert.Equal(t, r.CallRoute("/aaaa/bar"), Handler("d"))
}
