package gotl

import (
	"testing"
)

func dump_q[T any](t *testing.T, q *Queue[T]) {
	t.Errorf("len(d): %d, cap(d): %d, front: %d", len(q.d), cap(q.d), q.front)
}

func TestQueueBasic(t *testing.T) {
	var q Queue[int]

	if !q.Empty() {
		t.Errorf("The queue should be empty now!")
	}

	q.Push(1)
	if q.Len() != 1 {
		t.Errorf("Queue length should be 1, got %d", q.Len())
	}

	if q.Front() != 1 {
		t.Errorf("Expected 1 on the front, got %d", q.Front())
	}

	q.Pop()
	if q.Len() != 0 {
		t.Errorf("Queue length should be 0, got %d", q.Len())
		dump_q(t, &q)
	}

}
