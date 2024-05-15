package lists

const ListsVersion string = "1.2.0"

// Interface for countable types
type Countable[T any] interface {
	Count() uint
}

// Interface for sliceable types
type Sliceable[T any] interface {
	ToSlice() []any
}

// Interface for limited lists
type Limitable[T any] interface {
	Capacity() uint
	Isfull() bool
}

// Interface for a Fifo list
type Fifo[T any] interface {
	Enqueue(x T)
	Dequeue() (T, error)
	IsEmpty() bool
	Peek() (T, error)
}

// Interface for a Lifo list
type Lifo[T any] interface {
	Push(x T)
	Pop() (T, error)
	IsEmpty() bool
	Peek() (T, error)
}

// Struct for a single link node
type snode[T any] struct {
	data T
	next *snode[T]
}

// Constructor for a single link node
func newSingleNode[T any](d T, n *snode[T]) *snode[T] {
	return &snode[T]{
		data: d,
		next: n,
	}
}
