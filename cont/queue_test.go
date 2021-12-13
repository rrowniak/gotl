package cont

import (
	"testing"
)

type tmpStruct struct {
	a int
	b string
}

func dump_q[T any](t *testing.T, q *Queue[T]) {
	t.Errorf("len(d): %d, cap(d): %d, front: %d", len(q.d), cap(q.d), q.front)
}

func TestQueueExample(t *testing.T) {
	var q Queue[int]
	q.Push(1)
	q.Push(2)
	if q.Len() != 2 {
		t.Errorf("Expected queue len == 2, got %d", q.Len())
	}
	for i := 1; i < 3; i++ {
		if q.Front() != i {
			t.Errorf("Expected to have %d, got %d", i, q.Front())
		}
		q.Pop()
	}
	if !q.Empty() {
		t.Errorf("Queue should be empty now, but have %d element(s)", q.Len())
		dump_q(t, &q)
	}
}

func TestQueueEmptyPop(t *testing.T) {
	var q Queue[tmpStruct]

	if !q.Empty() {
		t.Errorf("Queue should be empty now, but have %d element(s)", q.Len())
		dump_q(t, &q)
	}

	q.Pop()

	if !q.Empty() {
		t.Errorf("Queue should be empty now, but have %d element(s)", q.Len())
		dump_q(t, &q)
	}
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

func TestReservationQueue(t *testing.T) {
	var q Queue[string]
	q.Reserve(1)
	if q.Len() != 0 {
		t.Errorf("Expected len == 0,  got %d", q.Len())
	}

	q.Push("a")
	q.Push("b")

	if q.Len() != 2 {
		t.Errorf("Expected len == 2,  got %d", q.Len())
	}

	q.Reserve(20)

	if q.Len() != 2 {
		t.Errorf("Expected len == 2,  got %d", q.Len())
	}

	q.Push("c")
	if q.Len() != 3 {
		t.Errorf("Expected len == 3,  got %d", q.Len())
	}

	res := []string{"a", "b", "c"}
	for i, r := range res {
		if q.Front() != r {
			t.Errorf("Expected #%d value %s,  got %s", i, r, q.Front())
		}
		q.Pop()
	}
}

func TestAutoShrink(t *testing.T) {
	var q Queue[int]
	for i := 0; i < 100; i++ {
		q.Push(i)
	}

	for i := 0; i < 100; i++ {
		if q.Front() != i {
			t.Errorf("Expected front == %d, got %d", i, q.Front())
		}
		q.Pop()
	}

	if !q.Empty() {
		t.Errorf("Queue should be empty now, but have %d element(s)", q.Len())
		dump_q(t, &q)
	}
}
