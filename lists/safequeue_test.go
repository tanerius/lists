package lists

import (
	"testing"
)

func BenchmarkRegularSafeQueueEnqueue(b *testing.B) {
	queue := NewSafeQueue[int]()

	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkRegularSafeQueueDequeue(b *testing.B) {
	queue := NewSafeQueue[int]()
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}

/*
func BenchmarkLimitedSizeSafeQueueEnqueue(b *testing.B) {
	queue := NewSafeQueue[int](10)

	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

func BenchmarkLimitedSizeSafeQueueDequeue(b *testing.B) {
	queue := NewSafeQueue[int](10)
	b.StopTimer()
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Dequeue()
	}
}
*/
