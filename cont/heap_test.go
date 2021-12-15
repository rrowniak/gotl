package cont

import (
	"math/rand"
	"testing"
)

func TLessInt(l, r int) bool {
	return l < r
}

func TestHeapExample(t *testing.T) {
	h := NewHeap[int](10, TLessInt)
	h.Push(1)
	h.Push(5)
	h.Push(2)

	if h.Pop() != 5 {
		t.Errorf("Expected max = 5, got %d", h.GetMax())
	}

	if h.Pop() != 2 {
		t.Errorf("Expected max = 2, got %d", h.GetMax())
	}

	if h.Pop() != 1 {
		t.Errorf("Expected max = 1, got %d", h.GetMax())
	}
}

func TestHeapPush(t *testing.T) {
	h := NewHeap(10, TLessInt)

	if !h.Empty() {
		t.Error("Expected empty heap")
	}

	max := -1

	for i := 1; i <= 30; i++ {
		v := rand.Intn(100)
		if v > max {
			max = v
		}
		h.Push(v)

		if h.Len() != i {
			t.Errorf("Expected heap size = %d, got %d", i, h.Len())
		}

		if h.GetMax() != max {
			t.Errorf("Max is expected to be %d, got %d", max, h.GetMax())
		}
	}
}

func TestHeapPop(t *testing.T) {
	h := NewHeap(10, TLessInt)
	SIZE := 100
	for i := 1; i <= SIZE; i++ {
		h.Push(i)
	}

	for i := SIZE - 1; i >= 0; i-- {
		h.Pop()

		if h.Len() != i {
			t.Errorf("Expected %d elements, got %d elements", i, h.Len())
		}

		if i != 0 && h.GetMax() != i {
			t.Errorf("Max is expected to be %d, got %d", i, h.GetMax())
			t.Errorf("Internal structure: %v", h.d)
		}
	}

	h.Pop()
}
