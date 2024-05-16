package lists_examples

import (
	"fmt"

	"github.com/tanerius/lists/lists"
)

func Examples() {
	// Create a regular queue of integers
	queue := lists.NewQueue[int]()

	for i := 0; i < 998; i++ {
		queue.Enqueue(i)
	}
	// enqueue elements
	queue.Enqueue(100)
	queue.Enqueue(101)
	queue.Enqueue(102)
	queue.Enqueue(103)

	// dequeue and print
	for i := 0; i < 998; i++ {
		queue.Dequeue()
	}

	fmt.Println(queue.ToSlice())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	/*
	   Output:
	       2 <nil>
	       3 <nil>
	       4 <nil>
	       0 empty list
	*/

	// Create a fixed size queue of 3 integers
	fixedQueue := lists.NewLSQueue[int](4)
	// enqueue elements
	fixedQueue.Enqueue(5)
	fixedQueue.Enqueue(6)
	fixedQueue.Enqueue(7)
	fixedQueue.Enqueue(8)
	fixedQueue.Enqueue(9)

	fmt.Println(fixedQueue.ToSlice())
	// dequeue and print
	fmt.Println(fixedQueue.Dequeue())
	fmt.Println(fixedQueue.Dequeue())
	fmt.Println(fixedQueue.Dequeue())
	fmt.Println(fixedQueue.Dequeue())

	stack := lists.NewStack[int]()

	for i := 0; i < 998; i++ {
		stack.Push(i)
	}

	stack.Push(105)
	stack.Push(106)
	stack.Push(107)
	stack.Push(108)

	for i := 0; i < 998; i++ {
		stack.Pop()
	}

	fmt.Println(stack.ToSlice())

	// enqueue elements
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

}
