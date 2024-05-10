package lists

import (
	"testing"
)

func BenchmarkStackPush(b *testing.B) {
	stack := NewStack[int]()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStackPop(b *testing.B) {
	stack := NewStack[int]()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
