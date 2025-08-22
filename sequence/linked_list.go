package sequence

import "errors"

var (
	ErrLinkedListEmpty = errors.New("empty")
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

// Len returns the number of elements in the sequence
func (l *LinkedList[T]) Len() int {
	return l.size
}

// Insert inserts to head
func (l *LinkedList[T]) Insert(v T) {
	node := NewNode[T](v)
	if l.head == nil {
		l.head = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}
	l.size += 1
}

// Peek peeks head
func (l *LinkedList[T]) Peek() (*T, error) {
	if l.head == nil {
		return nil, ErrLinkedListEmpty
	}
	return &l.head.v, nil
}

// Delete deletes the head
func (l *LinkedList[T]) Delete() (*T, error) {
	if l.head == nil {
		return nil, ErrLinkedListEmpty
	}
	node := l.head
	l.head = node.next
	if l.head != nil {
		l.head.prev = nil
	}
	l.size -= 1
	return &node.v, nil
}
