package lists

import (
	"errors"
	"sync"
)

// The SafeQueue is a collection of entities that are maintained in a sequence and can be modified
// by the addition of entities at one end of the sequence and the removal of entities from the
// other end of the sequence.
//
// # SafeQueue is threadsafe
//
// # Please note that while the type itself is safe the data should handle its own thread safety
//
// Queue is a list that implements the Fifo interface
type SafeQueue[T any] struct {
	mu          sync.RWMutex
	maxBuffSize uint
	curBuffSize uint
	head        *snode[T]
	tail        *snode[T]
}

// The constructor for a new Queue instance with elements of type T.
// If size > 0 then the generated queue is a limited size queue.
// If size = 0 then it is a generic unlimited queue
//
// Returns a pointer to a queue
func NewSafeQueue[T any](size uint) *SafeQueue[T] {
	return &SafeQueue[T]{
		maxBuffSize: size,
		curBuffSize: 0,
		head:        nil,
		tail:        nil,
	}
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *SafeQueue[T]) Enqueue(element T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.curBuffSize == 0 {
		r.head = newSingleNode[T](element, nil)
		r.tail = r.head
		r.curBuffSize++
	} else {
		if r.maxBuffSize > 0 && r.curBuffSize == r.maxBuffSize {
			r.head = r.head.next
		} else {
			r.curBuffSize++
		}
		newItem := newSingleNode(element, nil)
		r.tail.next = newItem
		r.tail = newItem
	}
}

// Remove and return am element of type T from the beginning of the queue. Complexity is O(1)
func (r *SafeQueue[T]) Dequeue() (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	ret := r.head.data
	r.head = r.head.next
	r.curBuffSize--
	if r.curBuffSize == 0 {
		r.head = nil
		r.tail = nil
	}
	return ret, nil
}

// Checks if the queue is empty
//
// Return true if empty false otherwise
func (r *SafeQueue[T]) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.curBuffSize == 0
}

// Checks if the limited size queue is full
//
// Return true if the limites size queue has reached its given capacity. Otherwise returns false.
// Also returns false if the queue was created with a size of 0 (generic unlimited queue)
func (r *SafeQueue[T]) Isfull() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.curBuffSize == r.maxBuffSize
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *SafeQueue[T]) Peek() (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	return r.tail.data, nil
}

// Return a slice representation of the current state of the queue
func (r *SafeQueue[T]) ToSlice() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.curBuffSize == 0 {
		return make([]T, 0)
	}

	s := make([]T, 0, r.curBuffSize)
	tmp := r.head

	for tmp != nil {
		s = append(s, tmp.data)
		tmp = tmp.next
	}
	return s
}

// Return the number of elements in the queue
func (r *SafeQueue[T]) Count() uint {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.curBuffSize
}
