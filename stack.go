package gotl

// Stack container
type Stack[T any] struct {
	d []T
}

// Increase the capacity of the stack. If the stack is not empty, the elements will be copied to the new location.
func (s *Stack[T]) Reserve(size int) {
	if s.Empty() {
		s.d = make([]T, 0, size)
	} else {
		new := make([]T, len(s.d), size)
		copy(new, s.d)
		s.d = new
	}
}

// Push an element.
func (s *Stack[T]) Push(t T) {
	s.d = append(s.d, t)
}

// Get an element from the top. This function does not change state of the stack.
// If this function is being called on an empty stack, the panic will be raised.
func (s Stack[T]) Top() T {
	return s.d[len(s.d)-1]
}

// Pop an element.
func (s *Stack[T]) Pop() {
	s.d = s.d[:len(s.d)-1]
}

// Return true if the stack is empty.
func (s Stack[t]) Empty() bool {
	return s.Len() == 0
}

// Get the amount of elements in the stack.
func (s Stack[t]) Len() int {
	return len(s.d)
}
