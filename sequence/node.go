package sequence

type Node[T comparable] struct {
	v    T
	next *Node[T]
	prev *Node[T]
}

func NewNode[T comparable](v T) *Node[T] {
	return &Node[T]{
		v:    v,
		next: nil,
		prev: nil,
	}
}
