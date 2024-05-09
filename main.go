package main

import (
	"fmt"

	"github.com/tanerius/lists/lists"
)

func main() {
	// create a limimted size queue
	testq := lists.NewQueue[int](5)

	testq.Enqueue(5)
	testq.Enqueue(6)
	testq.Enqueue(7)
	testq.Enqueue(8)
	testq.Enqueue(9)
	testq.Enqueue(10)

	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())

	/*
		OUTPUT:

		6 <nil>
		7 <nil>
		8 <nil>
		9 <nil>
		10 <nil>
		0 empty list
	*/

	tests := lists.NewStack[int]()
	tests.Push(1)
	tests.Push(2)
	tests.Push(3)
	tests.Push(4)
	fmt.Println(tests.Pop())
	fmt.Println(tests.Pop())
	fmt.Println(tests.Pop())
	fmt.Println(tests.Pop())
	fmt.Println(tests.Pop())

}
