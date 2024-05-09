package main

import (
	"fmt"

	"github.com/tanerius/lists/lists"
)

func main() {
	testq := lists.NewQueue[int](5)
	testq.Enqueue(5)
	testq.Enqueue(6)
	testq.Enqueue(7)

	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())
	fmt.Println(testq.Dequeue())

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
