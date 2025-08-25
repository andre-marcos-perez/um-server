package sequence

import "errors"

var (
	ErrSequenceEmpty = errors.New("empty sequence")
)

// Sequence defines an interface for sequence data structures.
// Ref.: https://ocw.mit.edu/courses/6-006-introduction-to-algorithms-spring-2020/resources/mit6_006s20_lec2/
type Sequence[T comparable] interface {
	Len() int
	Iter() []T
	Peek() (*T, error)
	Sort()
	Insert(v T)
	Delete() (*T, error)
}
