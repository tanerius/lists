package lists

import "errors"

// A Stack is an abstract data type that serves as a collection of elements with two main operations:
//
// Push, which adds an element to the collection, and
// Pop, which removes the most recently added element.
// Additionally, a peek operation can, without modifying the stack, return the value of the last element added.
type Stack[T any] struct {
	curBuffSize uint
	index       uint
	head        *arrnode[T]
}

// Constructs a new Stack with elements of type T
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{
		curBuffSize: 0,
		head:        newArrayNode[T](nil),
		index:       999,
	}
}

// Pushes a new element T onto the stack. Complexity is O(1)
func (r *Stack[T]) Push(element T) {
	if r.curBuffSize > 0 {
		if r.index == 0 {
			r.index = 999
			newNode := newArrayNode[T](r.head)
			r.head = newNode
		} else {
			r.index--
		}
	}
	r.curBuffSize++
	r.head.write(element, int(r.index))
}

// Removes the most recently added element T from the stack and returns it. Complexity is O(1)
func (r *Stack[T]) Pop() (T, error) {
	var result T
	if r.curBuffSize == 0 {
		return result, errors.New("empty list")
	}

	r.curBuffSize--
	result = r.head.read(int(r.index))
	r.index++

	if r.index > 999 {
		r.index = 0
		if r.head.next != nil {
			r.head = r.head.next
		}
	}

	return result, nil
}

// Checks to see if the stack is empty.
//
// Returns true if stack is empty otherwise false
func (r *Stack[T]) IsEmpty() bool {
	return r.curBuffSize == 0
}

// The Peek operation returns, without modifying the stack, the value of the last element T added
func (r *Stack[T]) Peek() (T, error) {
	var result T
	if r.curBuffSize == 0 {
		return result, errors.New("empty list")
	}

	return r.head.read(int(r.index)), nil
}

// Return a slice representation of the current state of the stack
func (r *Stack[T]) ToSlice() []T {
	s := make([]T, 0, r.curBuffSize)

	if r.curBuffSize == 0 {
		return s
	}

	tmp := r.head
	index := r.index
	finished := false
	cnt := 0

	for tmp != nil && !finished {
		s = append(s, tmp.read(int(index)))
		cnt++

		if index == 999 {
			index = 0
			tmp = tmp.next
		} else {
			index++
		}

		if cnt > int(r.curBuffSize) {
			finished = true
		}
	}
	return s
}

// Return the number of elements in the Stack
func (r *Stack[T]) Count() uint {
	return r.curBuffSize
}
