package lists

// Interface for a Fifo list
type Fifo interface {
	Enqueue(x any)
	Dequeue() (any, error)
	IsEmpty() bool
	Isfull() bool
	Peek() (any, error)
}

// Interface for a Lifo list
type Lifo interface {
	Push(x any)
	Pop() (any, error)
	IsEmpty() bool
	Peek() (any, error)
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
