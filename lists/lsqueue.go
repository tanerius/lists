package lists

import "errors"

// The LSQueue is a collection of entities that are maintained in a sequence and can be modified
// by the addition of entities at one end of the sequence and the removal of entities from the
// other end of the sequence. This is a specialised version of queue that has a limited size.
// When the limit is reached the first equeued element will be replaced by the next incoming enque.
//
// LSQueue is a list that implements the Fifo interface
type LSQueue[T any] struct {
	maxBuffSize uint
	curBuffSize uint
	lastIndex   int
	data        []T
}

// The constructor for a new LSQueue instance with elements of type T.
//
// # Returns a pointer to a LSQueue
func NewLSQueue[T any](size uint) *LSQueue[T] {

	return &LSQueue[T]{
		maxBuffSize: size,
		curBuffSize: 0,
		lastIndex:   0,
		data:        make([]T, size, size),
	}
}

func (r *LSQueue[T]) getFrontElementIndex() int {
	if r.curBuffSize == 0 {
		return 0
	}

	i := (r.lastIndex + 1) - int(r.curBuffSize)

	if i < 0 {
		i = int(r.maxBuffSize) + i
	}

	return i
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *LSQueue[T]) Enqueue(element T) {
	if r.maxBuffSize == 0 {
		return
	}

	if r.curBuffSize+1 < r.maxBuffSize {
		r.curBuffSize++
	}

	r.lastIndex = (r.lastIndex + 1) % int(r.maxBuffSize)
	r.data[r.lastIndex] = element
}

// Remove and return am element of type T from the beginning of the queue. Complexity is O(1)
func (r *LSQueue[T]) Dequeue() (T, error) {
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	indexOfElementToDequeue := r.getFrontElementIndex()

	r.curBuffSize--
	return r.data[indexOfElementToDequeue], nil
}

// Checks if the queue is empty
//
// Return true if empty false otherwise
func (r *LSQueue[T]) IsEmpty() bool {
	return r.curBuffSize == 0
}

// Checks if the limited size queue is full
//
// Return true if the limites size queue has reached its given capacity. Otherwise returns false.
// Also returns false if the queue was created with a size of 0 (generic unlimited queue)
func (r *LSQueue[T]) Isfull() bool {
	return r.curBuffSize == r.maxBuffSize
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *LSQueue[T]) Peek() (T, error) {
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	indexOfElementToDequeue := r.getFrontElementIndex()
	return r.data[indexOfElementToDequeue], nil
}

// Return a slice representation of the current state of the queue
func (r *LSQueue[T]) ToSlice() []T {
	if r.curBuffSize == 0 {
		return make([]T, 0)
	}

	s := make([]T, 0, r.curBuffSize)

	// TODO: Implement

	return s
}

// Return the number of elements in the queue
func (r *LSQueue[T]) Count() uint {
	return r.curBuffSize
}

// Return the number of elements in the queue
func (r *LSQueue[T]) Capacity() uint {
	return r.maxBuffSize
}
