package lists

import "errors"

// The Queue is a collection of entities that are maintained in a sequence and can be modified
// by the addition of entities at one end of the sequence and the removal of entities from the
// other end of the sequence.
//
// Queue is a list that implements the Fifo interface
type Queue[T any] struct {
	curBuffSize      uint
	lowClusterIndex  uint
	highClusterIndex uint
	lastElementIndex uint
	data             map[uint][1000]T
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
		lastElementIndex: 0,
		data:             make(map[uint][1000]T),
	}
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *Queue[T]) Enqueue(element T) {
	if r.curBuffSize == 0 {
		r.data[0] = [1000]T{}
	} else {
		r.lastElementIndex++
		clusterId := r.lastElementIndex / 1000
		normalizedClusterIndex := r.lastElementIndex % 1000
		// check if cluster exists
		if clusterId > r.highClusterIndex {
			// we need to create a new cluster here
			r.highClusterIndex++
			r.data[r.highClusterIndex] = [1000]T{}

			if clusterId != r.highClusterIndex {
				panic("misaligned cluster cannot continue")
			}
		}
		arr := r.data[clusterId]
		arr[normalizedClusterIndex] = element
	}
	r.curBuffSize++
}

// Remove and return am element of type T from the beginning of the queue. Complexity is O(1)
func (r *Queue[T]) Dequeue() (T, error) {
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
func (r *Queue[T]) IsEmpty() bool {
	return r.curBuffSize == 0
}

// Return am element of type T from the beginning of the queue without Dequeuing it. Complexity is O(1)
func (r *Queue[T]) Peek() (T, error) {
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	return r.tail.data, nil
}

// Return a slice representation of the current state of the queue
func (r *Queue[T]) ToSlice() []T {
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
func (r *Queue[T]) Count() uint {
	return r.curBuffSize
}
