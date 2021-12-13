package cont

import (
	"testing"
)

func TestStackExample(t *testing.T) {
	var s Stack[string]
	s.Push("bottom")
	s.Push("top")
	if s.Top() != "top" {
		t.Errorf("Expected 'top', got '%s'", s.Top())
	}
	s.Pop()
	if s.Top() != "bottom" {
		t.Errorf("Expected 'bottom', got '%s'", s.Top())
	}
	s.Pop()
	if !s.Empty() {
		t.Error("Stack is expected to be empty")
	}
}

func TestEmptyStack(t *testing.T) {
	var s Stack[float64]
	if !s.Empty() {
		t.Error("Stack is expected to be empty")
	}

	s.Push(1.0)
	if s.Empty() {
		t.Error("Stack is expected to be non-empty")
	}

	s.Pop()
	if !s.Empty() {
		t.Error("Stack is expected to be empty")
	}

	s.Push(2.0)
	if s.Empty() {
		t.Error("Stack is expected to be non-empty")
	}
}

func TestBasicStack(t *testing.T) {
	var s Stack[int]
	s.Push(1)
	s.Push(2)

	if s.Len() != 2 {
		t.Errorf("Expected len == 1,  got %d", s.Len())
	}

	top := s.Top()
	if top != 2 {
		t.Errorf("Expected top == 2, got %d", top)
	}

	s.Pop()
	s.Pop()

	if !s.Empty() {
		t.Errorf("Expected empty stack")
	}
}

func TestReservationStack(t *testing.T) {
	var s Stack[int]
	s.Reserve(1)
	if s.Len() != 0 {
		t.Errorf("Expected len == 0,  got %d", s.Len())
	}

	s.Push(10)
	s.Push(11)

	if s.Len() != 2 {
		t.Errorf("Expected len == 2,  got %d", s.Len())
	}

	s.Reserve(20)

	if s.Len() != 2 {
		t.Errorf("Expected len == 2,  got %d", s.Len())
	}

	s.Push(12)
	if s.Len() != 3 {
		t.Errorf("Expected len == 3,  got %d", s.Len())
	}

	res := []int{12, 11, 10}
	for i, r := range res {
		if s.Top() != r {
			t.Errorf("Expected #%d value %d,  got %d", i, r, s.Top())
		}
		s.Pop()
	}
}
