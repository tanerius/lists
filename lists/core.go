package lists

const ListsVersion string = "1.1.1"

// Interface for countable types
type Countable interface {
	Count() uint
}

// Interface for sliceable types
type Sliceable interface {
	ToSlice() []any
}

// Interface for a Fifo list
type Fifo[T any] interface {
	Enqueue(x T)
	Dequeue() (T, error)
	IsEmpty() bool
	Isfull() bool
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
