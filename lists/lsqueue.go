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
	head        *snode[T]
	tail        *snode[T]
}

// The constructor for a new LSQueue instance with elements of type T.
//
// # Returns a pointer to a LSQueue
func NewLSQueue[T any](size uint) *LSQueue[T] {
	return &LSQueue[T]{
		maxBuffSize: size,
		curBuffSize: 0,
		head:        nil,
		tail:        nil,
	}
}

// Add an element of type T to the end of the queue. Complexity is O(1)
func (r *LSQueue[T]) Enqueue(element T) {
	if r.maxBuffSize == 0 {
		return
	}

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
func (r *LSQueue[T]) Dequeue() (T, error) {
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

	return r.tail.data, nil
}

// Return a slice representation of the current state of the queue
func (r *LSQueue[T]) ToSlice() []T {
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
func (r *LSQueue[T]) Count() uint {
	return r.curBuffSize
}

// Return the number of elements in the queue
func (r *LSQueue[T]) Capacity() uint {
	return r.maxBuffSize
}
