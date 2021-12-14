package cont

import (
	"testing"
)

func TestPriorityQueueExample(t *testing.T) {
	q := NewPriorityQueue(100, TLessInt)
	q.Push(1)
	q.Push(1)
	q.Push(2)

	if q.Front() != 2 {
		t.Errorf("Front expected to be 2, got %d", q.Front())
	}
	q.Pop()

	if q.Front() != 1 {
		t.Errorf("Front expected to be 1, got %d", q.Front())
	}
	q.Pop()

	if q.Front() != 1 {
		t.Errorf("Front expected to be 1, got %d", q.Front())
	}
	q.Pop()
}

func TestPriorityQueueBasicScenario(t *testing.T) {
	q := NewPriorityQueue(100, TLessInt)

	if !q.Empty() {
		t.Error("Priority queue is expected to be empty")
	}

	for i := 1; i <= 10; i++ {
		q.Push(1)
	}

	q.Push(2)

	if q.Front() != 2 {
		t.Errorf("Front expected to be 2, got %d", q.Front())
	}

	if q.Len() != 11 {
		t.Errorf("Front expected 11 elements, got %d", q.Len())
	}

	q.Pop()

	for i := 1; i <= 10; i++ {
		if q.Front() != 1 {
			t.Errorf("Front expected to be 1, got %d", q.Front())
		}
		q.Pop()
	}

}

type PQTypeCust struct {
	a int
	b string
}

func custTypeLess(a, b PQTypeCust) bool {
	return a.a < b.a
}
func TestPriorityQueueCustomType(t *testing.T) {
	pq := NewPriorityQueue[PQTypeCust](10, custTypeLess)
	pq.Push(PQTypeCust{1, "a"})
	pq.Push(PQTypeCust{2, "b"})
	pq.Push(PQTypeCust{1, "c"})
	pq.Push(PQTypeCust{-1, "ddd"})

	if pq.Front().a != 2 {
		t.Errorf("Expected {2, b}, got {%d, %s}", pq.Front().a, pq.Front().b)
	}
}
