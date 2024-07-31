package main

import (
	"slices"
	"sort"
	"strings"
)

/*

We want to implement a middleware router for our web service,
which based on the path returns different strings (these would represent “functions to invoke” in a real application).

Our interface for the router looks something like:


interface Router {
  fun addRoute(path: String, result: String) : Unit;
  fun callRoute(path: String) : String;
}

Usage:
Router.addRoute("/bar", "result")
Router.callRoute("/bar") -> "result"

Although do note that you do not have to implement this interface directly as written,
there are different ways of abstracting this behaviour that you are free to explore.
The important part is that your router has a way to declare routes, and then act upon a path.
*/

/*
james notes: I'm going straight to scale ups, which wants you to pass path paramaters like:
"Allow the ability to add arguments to the handlers via path parameters."
The scale up before this is for * wild card matching, I'll start with that.
*/
/*
My first thoughts were to construct a tree, like a file directory tree.
Which I think is still worst case O(n) but still improves runtime.
This seemed pretty hard to do in Golang. But it could be done.

The way the solution does it is to just use an array for the pattern matching.
Fairly slow if you have a lot of routes.
I think it's fairly trivial using a list, so I'm gonna go hard mode and try a tree
*/

type Handler string

type Router struct {
	routes Tree
}

func NewRouter() *Router {
	return &Router{routes: *NewTree()}
}

func (r *Router) AddRoute(path string, handler Handler) {
	pathArr := strings.Split(path, "/")
	slices.Insert(pathArr, 0, "/")
	r.routes.Put(pathArr, handler)
}

func (r *Router) CallRoute(path string) Handler {
	pathArr := strings.Split(path, "/")
	slices.Insert(pathArr, 0, "/")
	return r.routes.Get(pathArr)
}

type Tree struct {
	key      string
	handler  Handler
	subtrees []*Tree
}

func NewTree() *Tree {
	return &Tree{key: "/", subtrees: make([]*Tree, 0)}
}

// Put overwrites the handler at a given path, or creates a new
// subtree with the given path if it does not exist.
func (t *Tree) Put(path []string, handler Handler) {
	if len(path) == 1 {
		t.handler = handler
		t.key = path[0]
		return
	}

	for _, subtree := range t.subtrees {
		if subtree.key == path[1] {
			subtree.Put(path[1:], handler)
			return
		}
	}
	newTree := &Tree{}
	newTree.key = path[1]
	newTree.Put(path[1:], handler)
	t.subtrees = append(t.subtrees, newTree)
	sort.Slice(t.subtrees, func(i, j int) bool {
		return t.subtrees[i].key > t.subtrees[j].key
	})
}

// Get returns the handler at given path, or a "not found" handler if not found
func (t *Tree) Get(path []string) Handler {
	if len(path) == 1 {
		return t.handler
	}
	for _, subtree := range t.subtrees {
		if subtree.key == path[1] || subtree.key == "*" {
			return subtree.Get(path[1:])
		}
	}
	return "404 not found"
}
