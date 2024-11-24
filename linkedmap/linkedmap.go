package linkedmap

import "fmt"

// Entry represents a key-value pair in the LinkedMap.
type Entry[K comparable, V any] struct {
	key   K
	value V
}

// Node represents a node in the doubly linked list.
type Node[K comparable, V any] struct {
	value *Entry[K, V]
	prev  *Node[K, V]
	next  *Node[K, V]
}

// LinkedMap is a combination of a map and a doubly-linked list to maintain insertion order.
type LinkedMap[K comparable, V any] struct {
	capacity int
	data     map[K]*Node[K, V]
	head     *Node[K, V]
	tail     *Node[K, V]
}

// NewLinkedMap creates a new LinkedMap with the given capacity.
func NewLinkedMap[K comparable, V any](initialCapacity int) *LinkedMap[K, V] {
	return &LinkedMap[K, V]{
		data: make(map[K]*Node[K, V], initialCapacity),
	}
}

// Get retrieves the value for a key in the map, updating its recentness.
func (lm *LinkedMap[K, V]) Get(key K) (V, bool) {
	var zeroValue V
	if node, found := lm.data[key]; found {
		lm.moveToFront(node)
		return node.value.value, true
	}
	return zeroValue, false
}

func (lm *LinkedMap[K, V]) GetNewest() (*Entry[K, V], bool) {
	if lm.head != nil {
		return lm.head.value, true
	}
	var e *Entry[K, V]
	return e, false
}

func (lm *LinkedMap[K, V]) GetOldest() (*Entry[K, V], bool) {
	if lm.tail != nil {
		return lm.tail.value, true
	}
	var e *Entry[K, V]
	return e, false
}

// Put inserts a key-value pair into the map, updating its recentness.
func (lm *LinkedMap[K, V]) Put(key K, value V) {
	if node, found := lm.data[key]; found {
		node.value.value = value
		lm.moveToFront(node)
	} else {
		newNode := &Node[K, V]{value: &Entry[K, V]{key, value}}
		lm.addToFront(newNode)
		lm.data[key] = newNode
	}
}

// Remove deletes a key from the map.
func (lm *LinkedMap[K, V]) Remove(key K) {
	if node, found := lm.data[key]; found {
		lm.removeNode(node)
		delete(lm.data, key)
	}
}

// moveToFront moves a node to the front of the linked list.
func (lm *LinkedMap[K, V]) moveToFront(node *Node[K, V]) {
	if lm.head == node {
		return
	}
	lm.removeNode(node)
	lm.addToFront(node)
}

// addToFront adds a node to the front of the linked list.
func (lm *LinkedMap[K, V]) addToFront(node *Node[K, V]) {
	if lm.head == nil {
		// First element
		lm.head = node
		lm.tail = node
	} else {
		node.next = lm.head
		lm.head.prev = node
		lm.head = node
	}
}

// removeNode removes a node from the linked list.
func (lm *LinkedMap[K, V]) removeNode(node *Node[K, V]) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		lm.head = node.next
	}

	if node.next != nil {
		node.next.prev = node.prev
	} else {
		lm.tail = node.prev
	}
}

// removeOldest removes the oldest node from the linked list.
func (lm *LinkedMap[K, V]) removeOldest() {
	if lm.tail != nil {
		delete(lm.data, lm.tail.value.key)
		lm.removeNode(lm.tail)
	}
}

// Print prints all entries in the map from most to least recently used.
func (lm *LinkedMap[K, V]) Print() {
	for node := lm.head; node != nil; node = node.next {
		fmt.Printf("%v: %v\n", node.value.key, node.value.value)
	}
}
