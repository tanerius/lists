package lists

import (
	"errors"
)

// The Queue is a collection of entities that are maintained in a sequence and can be modified
// by the addition of entities at one end of the sequence and the removal of entities from the
// other end of the sequence.
//
// Queue is a list that implements the Fifo interface
type Queue[T any] struct {
	curBuffSize uint
	headIndex   uint
	tailIndex   uint
	head        *arrnode[T]
	tail        *arrnode[T]
}

// The constructor for a new Queue instance with elements of type T.
// If size > 0 then the generated queue is a limited size queue.
// If size = 0 then it is a generic unlimited queue
//
// Returns a pointer to a queue
func NewQueue[T any]() *Queue[T] {
	node := newArrayNode[T](nil)
	return &Queue[T]{
		curBuffSize: 0,
		headIndex:   0,
		tailIndex:   0,
		head:        node,
		tail:        node,
	}
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *Queue[T]) Enqueue(element T) {
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
func (r *Queue[T]) Dequeue() (T, error) {
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
func (r *Queue[T]) IsEmpty() bool {
	return r.curBuffSize == 0
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *Queue[T]) Peek() (T, error) {
	var result T
	if r.curBuffSize == 0 {
		return result, errors.New("empty list")
	}

	normalizedIndex := 1000 % r.headIndex
	result = r.head.read(int(normalizedIndex))

	return result, nil
}

// Return a slice representation of the current state of the queue
func (r *Queue[T]) ToSlice() []T {
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
func (r *Queue[T]) Count() uint {
	return r.curBuffSize
}
