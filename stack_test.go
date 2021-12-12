package gotl

import (
	"testing"
)

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
