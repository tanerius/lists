package lists

import "errors"

// A Stack is an abstract data type that serves as a collection of elements with two main operations:
//
// Push, which adds an element to the collection, and
// Pop, which removes the most recently added element.
// Additionally, a peek operation can, without modifying the stack, return the value of the last element added.
type Stack[T any] struct {
	curBuffSize int
	head        *snode[T]
	tail        *snode[T]
}

// Constructs a new Stack with elements of type T
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		curBuffSize: 0,
		head:        nil,
		tail:        nil,
	}
}

// Pushes a new element T onto the stack. Complexity is O(1)
func (r *Stack[T]) Push(element T) {
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
func (r *Stack[T]) Pop() (T, error) {
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
func (r *Stack[T]) IsEmpty() bool {
	return r.curBuffSize == 0
}

// The Peek operation returns, without modifying the stack, the value of the last element T added
func (r *Stack[T]) Peek() (T, error) {
	if r.curBuffSize == 0 {
		var result T
		return result, errors.New("empty list")
	}

	return r.head.data, nil
}

func (r *Stack[T]) ToSlice() []T {
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
