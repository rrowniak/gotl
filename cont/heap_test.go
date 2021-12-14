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
	h.Insert(1)
	h.Insert(5)
	h.Insert(2)

	if h.GetMax() != 5 {
		t.Errorf("Expected max = 5, got %d", h.GetMax())
	}
	h.RemoveMax()

	if h.GetMax() != 2 {
		t.Errorf("Expected max = 2, got %d", h.GetMax())
	}
	h.RemoveMax()

	if h.GetMax() != 1 {
		t.Errorf("Expected max = 1, got %d", h.GetMax())
	}
	h.RemoveMax()
}

func TestHeapInsert(t *testing.T) {
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
		h.Insert(v)

		if h.Len() != i {
			t.Errorf("Expected heap size = %d, got %d", i, h.Len())
		}

		if h.GetMax() != max {
			t.Errorf("Max is expected to be %d, got %d", max, h.GetMax())
		}
	}
}

func TestHeapRemoveMax(t *testing.T) {
	h := NewHeap(10, TLessInt)
	SIZE := 100
	for i := 1; i <= SIZE; i++ {
		h.Insert(i)
	}

	for i := SIZE - 1; i >= 0; i-- {
		h.RemoveMax()

		if h.Len() != i {
			t.Errorf("Expected %d elements, got %d elements", i, h.Len())
		}

		if i != 0 && h.GetMax() != i {
			t.Errorf("Max is expected to be %d, got %d", i, h.GetMax())
			t.Errorf("Internal structure: %v", h.d)
		}
	}

	h.RemoveMax()
}
