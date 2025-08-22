package sequence

type Sequence[T comparable] interface {
	Len() int
	Peek() (*T, error)
	Insert(v T)
}
