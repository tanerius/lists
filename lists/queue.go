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
	curBuffSize      uint
	lowClusterIndex  uint
	highClusterIndex uint
	headIndex        uint
	tailIndex        uint
	data             map[uint]*[1000]T
}

// The constructor for a new Queue instance with elements of type T.
// If size > 0 then the generated queue is a limited size queue.
// If size = 0 then it is a generic unlimited queue
//
// Returns a pointer to a queue
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		curBuffSize:      0,
		lowClusterIndex:  0,
		highClusterIndex: 0,
		headIndex:        0,
		tailIndex:        0,
		data:             make(map[uint]*[1000]T),
	}
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *Queue[T]) Enqueue(element T) {
	if r.curBuffSize == 0 {
		r.data[0] = &[1000]T{}
		r.data[0][0] = element
	} else {
		r.tailIndex++
		clusterId := r.tailIndex / 1000
		normalizedClusterIndex := r.tailIndex % 1000
		// check if cluster exists
		if clusterId > r.highClusterIndex {
			// we need to create a new cluster here
			r.highClusterIndex++
			r.data[r.highClusterIndex] = &[1000]T{}

			if clusterId != r.highClusterIndex {
				panic("misaligned cluster cannot continue")
			}
		}
		r.data[clusterId][normalizedClusterIndex] = element
	}
	r.curBuffSize++
}

// Remove and return am element of type T from the beginning of the queue. Complexity is O(1)
func (r *Queue[T]) Dequeue() (T, error) {
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}
	var ret T
	clusterId := r.headIndex / 1000
	normalizedClusterIndex := r.headIndex % 1000
	if r.curBuffSize == 1 {
		ret = r.data[clusterId][normalizedClusterIndex]
		r.curBuffSize = 0
		r.headIndex = 0
		r.tailIndex = 0
		if clusterId > 0 {
			r.data = make(map[uint]*[1000]T)
		}
	} else {
		ret = r.data[clusterId][normalizedClusterIndex]
		r.headIndex++
		r.curBuffSize--
	}

	return ret, nil
}

// Checks if the queue is empty
//
// Return true if empty false otherwise
func (r *Queue[T]) IsEmpty() bool {
	return r.curBuffSize == 0
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *Queue[T]) Peek() (T, error) {
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	clusterId := r.headIndex / 1000
	normalizedClusterIndex := r.headIndex % 1000
	return r.data[clusterId][normalizedClusterIndex], nil

}

// Return a slice representation of the current state of the queue
func (r *Queue[T]) ToSlice() []T {

	return make([]T, 0)
}

// Return the number of elements in the queue
func (r *Queue[T]) Count() uint {
	return r.curBuffSize
}
