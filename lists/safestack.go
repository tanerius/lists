package lists

import (
	"errors"
	"sync"
)

// A SafeStack is an abstract data type that serves as a collection of elements with two main operations:
//
// # SafeStack is threadsafe
//
// # Please note that while the type itself is safe the data should handle its own thread safety
//
// Push, which adds an element to the collection, and
// Pop, which removes the most recently added element.
// Additionally, a peek operation can, without modifying the stack, return the value of the last element added.
type SafeStack[T any] struct {
	mu          sync.RWMutex
	curBuffSize uint
	head        *snode[T]
	tail        *snode[T]
}

// Constructs a new Stack with elements of type T
func NewSafeStack[T any]() *SafeStack[T] {
	return &SafeStack[T]{
		curBuffSize: 0,
		head:        nil,
		tail:        nil,
	}
}

// Pushes a new element T onto the stack. Complexity is O(1)
func (r *SafeStack[T]) Push(element T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.curBuffSize == 0 {
		r.head = newSingleNode[T](element, nil)
		r.tail = r.head
		r.curBuffSize++
	} else {
		r.curBuffSize++
		newItem := newSingleNode(element, r.head)
		r.head = newItem
	}
}

// Removes the most recently added element T from the stack and returns it. Complexity is O(1)
func (r *SafeStack[T]) Pop() (T, error) {
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

// Checks to see if the stack is empty.
//
// Returns true if stack is empty otherwise false
func (r *SafeStack[T]) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.curBuffSize == 0
}

// The Peek operation returns, without modifying the stack, the value of the last element T added
func (r *SafeStack[T]) Peek() (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	return r.head.data, nil
}

// Return a slice representation of the current state of the stack
func (r *SafeStack[T]) ToSlice() []T {
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

// Return the number of elements in the Stack
func (r *SafeStack[T]) Count() uint {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.curBuffSize
}
