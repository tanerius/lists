package lists

import (
	"testing"
)

func BenchmarkRegularQueueEnqueue(b *testing.B) {
	queue := NewQueue[int]()

	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkRegularQueueDequeue(b *testing.B) {
	queue := NewQueue[int]()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}

func BenchmarkLimitedSizeQueueEnqueue(b *testing.B) {
	queue := NewLSQueue[int](10)

	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkLimitedSizeQueueDequeue(b *testing.B) {
	queue := NewLSQueue[int](10)
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}
