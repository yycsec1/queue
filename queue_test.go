package queue

import (
	"testing"
	"time"
)

func TestQueue(t *testing.T) {
	q := New[string]()
	q.Push("John")
	q.Push("Jane")

	result1, ok1 := q.Pop()
	if result1 != "John" && ok1 != true {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result1, "John")
	}

	result2, ok2 := q.Peek()
	if result2 != "Jane" && ok2 != true {
		t.Errorf("Result was incorrect, got: %s, want: %s.", result2, "Jane")
	}

	result3 := q.Len()
	if result3 != 1 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result3, 1)
	}

	q.Pop()
	result4 := q.IsEmpty()
	if !result4 {
		t.Errorf("Result was incorrect, got: %t, want: %t.", result4, true)
	}

}

func TestPerformance(t *testing.T) {
	q := New[int]()
	const size = 10_000_000
	start := time.Now()
	for i := 0; i < size; i++ {
		q.Push(i)
	}
	elapsed := time.Since(start)
	t.Log("Time for 10 million Push() operation", elapsed)

	start = time.Now()
	for i := 0; i < size; i++ {
		q.Pop()
	}
	elapsed = time.Since(start)
	t.Log("Time for 10 million Pop() operation", elapsed)
}
