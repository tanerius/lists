package lists

import (
	"testing"
)

func TestQueue(t *testing.T) {
	queue := NewQueue[int]()

	// test if empty
	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", queue.IsEmpty(), true)
	}

	// test if full
	if queue.IsFull() {
		t.Errorf("IsFull() = %v, want %v", queue.IsFull(), false)
	}

	// check count method
	if queue.Count() != 0 {
		t.Errorf("Count() = %v, want %v", queue.Count(), 0)
	}

	// dequeue from an empty queue
	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Dequeue() = %v, want %v", err, "empty list")
	}

	// peek from an empty queue
	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Peek() = %v, want %v", err, "empty list")
	}

	// enqueue 1000 elements
	for i := 0; i < 1000; i++ {
		queue.Enqueue(i)
	}

	// test if empty
	if queue.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", queue.IsEmpty(), false)
	}

	// test if full
	if queue.IsFull() {
		t.Errorf("IsFull() = %v, want %v", queue.IsFull(), false)
	}

	// peak the first element
	element, err := queue.Peek()
	if element != 0 || err != nil {
		t.Errorf("Peek() = %v, %v, want %v, %v", element, err, 0, nil)
	}

	// check Count method
	if queue.Count() != 1000 {
		t.Errorf("Count() = %v, want %v", queue.Count(), 1000)
	}

	// dequeue 1000 elements
	for i := 0; i < 1000; i++ {
		element, _ := queue.Dequeue()
		if element != i {
			t.Errorf("Dequeue() = %v, want %v", element, i)
		}
	}
}

func TestLSQueue(t *testing.T) {
	queue := NewLSQueue[int](100)

	// test if empty
	if !queue.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", queue.IsEmpty(), true)
	}

	// test if full
	if queue.IsFull() {
		t.Errorf("IsFull() = %v, want %v", queue.IsFull(), false)
	}

	// check count method
	if queue.Count() != 0 {
		t.Errorf("Count() = %v, want %v", queue.Count(), 0)
	}

	//check Capacity method
	if queue.Capacity() != 100 {
		t.Errorf("Capacity() = %v, want %v", queue.Capacity(), 100)
	}

	// dequeue from an empty queue
	_, err := queue.Dequeue()
	if err == nil {
		t.Errorf("Dequeue() = %v, want %v", err, "empty list")
	}

	// peek from an empty queue
	_, err = queue.Peek()
	if err == nil {
		t.Errorf("Peek() = %v, want %v", err, "empty list")
	}

	// enqueue 1000 elements
	for i := 0; i < 1000; i++ {
		queue.Enqueue(i)
	}

	// test if empty
	if queue.IsEmpty() {
		t.Errorf("IsEmpty() = %v, want %v", queue.IsEmpty(), false)
	}

	// test if full
	if queue.IsFull() {
		t.Errorf("IsFull() = %v, want %v", queue.IsFull(), false)
	}

	// peak the first element
	element, err := queue.Peek()
	if element != 901 || err != nil {
		t.Errorf("Peek() = %v, %v, want %v, %v", element, err, 901, nil)
	}

	// check Count method
	if queue.Count() != 99 {
		t.Errorf("Count() = %v, want %v", queue.Count(), 99)
	}

	// dequeue 1000 elements
	for i := 0; i < 100; i++ {
		element, _ := queue.Dequeue()
		if !queue.IsEmpty() && element != i+901 {
			t.Errorf("Dequeue() = %v, want %v", element, i+901)
		}
	}
}

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
