package lists

import (
	"errors"
	"sync"
)

// The LSQueue is a collection of entities that are maintained in a sequence and can be modified
// by the addition of entities at one end of the sequence and the removal of entities from the
// other end of the sequence. This is a specialised version of queue that has a limited size.
// When the limit is reached the first equeued element will be replaced by the next incoming enque.
//
// SafeLSQueue is a thread safe version of LSQueue. However only the queue structure itself is safe. It is up to the
// developer to ensure thread safety of the internals of the data.
//
// SafeLSQueue is a list that implements the Fifo interface
type SafeLSQueue[T any] struct {
	maxBuffSize uint
	curBuffSize uint
	lastIndex   int
	data        []T
	mu          sync.RWMutex
}

// The constructor for a new LSQueue instance with elements of type T.
//
// # Returns a pointer to a LSQueue
func NewSafeLSQueue[T any](size uint) *SafeLSQueue[T] {

	return &SafeLSQueue[T]{
		maxBuffSize: size,
		curBuffSize: 0,
		lastIndex:   0,
		data:        make([]T, size),
	}
}

// A hidden method to compute the index of the next element to be dequeued
func (r *SafeLSQueue[T]) getFrontElementIndex() int {
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
func (r *SafeLSQueue[T]) Enqueue(element T) {
	r.mu.Lock()
	defer r.mu.Unlock()

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
func (r *SafeLSQueue[T]) Dequeue() (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

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
func (r *SafeLSQueue[T]) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.curBuffSize == 0
}

// Checks if the limited size queue is full
//
// Return true if the limites size queue has reached its given capacity. Otherwise returns false.
// Also returns false if the queue was created with a size of 0 (generic unlimited queue)
func (r *SafeLSQueue[T]) Isfull() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.curBuffSize == r.maxBuffSize
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *SafeLSQueue[T]) Peek() (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	indexOfElementToDequeue := r.getFrontElementIndex()
	return r.data[indexOfElementToDequeue], nil
}

// Return a slice representation of the current state of the queue
func (r *SafeLSQueue[T]) ToSlice() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.curBuffSize == 0 {
		return make([]T, 0)
	}
	indexOfInitialElement := r.getFrontElementIndex()
	s := make([]T, 0, r.curBuffSize)
	size := 0

	for size < int(r.curBuffSize) {
		s = append(s, r.data[indexOfInitialElement])
		indexOfInitialElement = (indexOfInitialElement + 1) % int(r.maxBuffSize)
		size++
	}

	return s
}

// Return the number of elements in the queue
func (r *SafeLSQueue[T]) Count() uint {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.curBuffSize
}

// Return the number of elements in the queue
func (r *SafeLSQueue[T]) Capacity() uint {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.maxBuffSize
}
