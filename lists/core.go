package lists

const ListsVersion string = "1.4.0"

// Interface for a Fifo list
type Fifo[T any] interface {
	Capacity() int
	Count() uint
	Dequeue() (T, error)
	Enqueue(x T)
	IsEmpty() bool
	IsFull() bool
	Peek() (T, error)
	ToSlice() []T
}

// Interface for a Lifo list
type Lifo[T any] interface {
	Count() uint
	IsEmpty() bool
	Push(x T)
	Pop() (T, error)
	Peek() (T, error)
	ToSlice() []T
}

// Struct for a single link node
type arrnode[T any] struct {
	data [1000]T
	next *arrnode[T]
}

// Constructor for a single link node
func newArrayNode[T any](n *arrnode[T]) *arrnode[T] {
	return &arrnode[T]{
		data: [1000]T{},
		next: n,
	}
}

func (r *arrnode[T]) write(value T, pos int) {
	r.data[pos] = value
}

func (r *arrnode[T]) read(pos int) T {
	return r.data[pos]
}
