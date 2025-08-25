package sequence

import (
	"math/rand"
	"time"
)

// LinkedList implements an unsorted double-linked list
type LinkedList[T comparable] struct {
	size int
	head *Node[T]
}

// NewLinkedList builds a new empty unsorted double-linked list
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		size: 0,
		head: nil,
	}
}

// Len returns the number of elements
func (list *LinkedList[T]) Len() int {
	return list.size
}

// Iter returns the elements in-order
func (list *LinkedList[T]) Iter() []T {
	elements := make([]T, 0, list.size)
	for node := list.head; node != nil; node = node.next {
		elements = append(elements, node.v)
	}
	return elements
}

// Peek peeks the element in the head
func (list *LinkedList[T]) Peek() (*T, error) {
	if list.head == nil {
		return nil, ErrSequenceEmpty
	}
	return &list.head.v, nil
}

// Sort randomly sorts the elements in place using Knuth shuffle algorithm
func (list *LinkedList[T]) Sort() {
	n := list.size
	nodes := make([]*Node[T], 0, n)
	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	// Collect the nodes in a slice of pointers
	for node := list.head; node != nil; node = node.next {
		nodes = append(nodes, node)
	}
	// Knuth shuffle algorithm
	for i := n - 1; i > 0; i -= 1 {
		j := random.Intn(i + 1)
		nodes[i].v, nodes[j].v = nodes[j].v, nodes[i].v
	}
}

// Insert inserts an element to head
func (list *LinkedList[T]) Insert(v T) {
	node := NewNode[T](v)
	if list.head == nil {
		list.head = node
	} else {
		node.next = list.head
		list.head.prev = node
		list.head = node
	}
	list.size += 1
}

// Delete deletes the element in the head
func (list *LinkedList[T]) Delete() (*T, error) {
	if list.head == nil {
		return nil, ErrSequenceEmpty
	}
	node := list.head
	list.head = node.next
	if list.head != nil {
		list.head.prev = nil
	}
	list.size -= 1
	return &node.v, nil
}
