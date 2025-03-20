package lists

import (
	"errors"
	"sync"
)

// The SafeQueue is a collection of entities that are maintained in a sequence and can be modified
// by the addition of entities at one end of the sequence and the removal of entities from the
// other end of the sequence.
//
// SafeQueue is a thread safe version of Queue. However only the queue structure itself is safe. It is up to the
// developer to ensure thread safety of the internals of the data.
//
// SafeQueue is a list that implements the Fifo interface
type SafeQueue[T any] struct {
	curBuffSize uint
	headIndex   uint
	tailIndex   uint
	head        *arrnode[T]
	tail        *arrnode[T]
	mu          sync.RWMutex
}

// The constructor for a new Queue instance with elements of type T.
// If size > 0 then the generated queue is a limited size queue.
// If size = 0 then it is a generic unlimited queue
//
// Returns a pointer to a queue
func NewSafeQueue[T any]() Fifo[T] {
	node := newArrayNode[T](nil)
	return &SafeQueue[T]{
		curBuffSize: 0,
		headIndex:   0,
		tailIndex:   0,
		head:        node,
		tail:        node,
	}
}

// Return the number of elements in the queue. -1 means unlimited
func (r *SafeQueue[T]) Capacity() int {
	return -1
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *SafeQueue[T]) Enqueue(element T) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.tail.write(element, int(r.tailIndex))

	if r.tailIndex == 999 {
		// prepare a new tail node for extra 1000 entries
		node := newArrayNode[T](nil)
		r.tail.next = node
		r.tail = r.tail.next
		r.tailIndex = 0
	} else {
		r.tailIndex++
	}
	r.curBuffSize++
}

// Remove and return am element of type T from the beginning of the queue. Complexity is O(1)
func (r *SafeQueue[T]) Dequeue() (T, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var result T
	if r.curBuffSize == 0 {
		return result, errors.New("empty list")
	}

	r.curBuffSize--
	result = r.head.read(int(r.headIndex))

	if r.headIndex == 999 {
		r.headIndex = 0
		r.head = r.head.next
	} else {
		r.headIndex++
	}

	if r.curBuffSize == 0 {
		r.headIndex = 0
		r.tailIndex = 0
	}

	return result, nil
}

// Checks if the queue is empty
//
// Return true if empty false otherwise
func (r *SafeQueue[T]) IsEmpty() bool {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.curBuffSize == 0
}

// Checks if the queue is full
func (r *SafeQueue[T]) IsFull() bool {
	return false
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *SafeQueue[T]) Peek() (T, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var result T
	if r.curBuffSize == 0 {
		return result, errors.New("empty list")
	}

	result = r.head.read(int(r.headIndex))

	return result, nil
}

// Return a slice representation of the current state of the queue
func (r *SafeQueue[T]) ToSlice() []T {
	r.mu.RLock()
	defer r.mu.RUnlock()

	ret := make([]T, 0)
	if r.curBuffSize == 0 {
		return ret
	}

	startingIndex := r.headIndex
	startingNode := r.head
	finished := false

	for startingNode != nil && !finished {
		ret = append(ret, startingNode.data[startingIndex])

		if r.tail == startingNode && (startingIndex+1) == r.tailIndex {
			return ret
		}

		if startingIndex == 999 {
			startingNode = startingNode.next
			startingIndex = 0
		} else {
			startingIndex++
		}
	}

	return ret
}

// Return the number of elements in the queue
func (r *SafeQueue[T]) Count() uint {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.curBuffSize
}
