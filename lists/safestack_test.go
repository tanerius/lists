package lists

import (
	"testing"
)

func BenchmarkSafeStackPush(b *testing.B) {
	stack := NewSafeStack[int]()

	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkSafeStackPop(b *testing.B) {
	stack := NewSafeStack[int]()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		stack.Pop()
	}
}
