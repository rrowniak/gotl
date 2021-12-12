package gotl

type Stack[T any] struct {
	d []T
}

func (s *Stack[T]) Reserve(size int) {
	if s.Empty() {
		s.d = make([]T, 0, size)
	} else {
		new := make([]T, len(s.d), size)
		copy(new, s.d)
		s.d = new
	}
}

func (s *Stack[T]) Push(t T) {
	s.d = append(s.d, t)
}

func (s Stack[T]) Top() T {
	return s.d[len(s.d)-1]
}

func (s *Stack[T]) Pop() {
	s.d = s.d[:len(s.d)-1]
}

func (s *Stack[t]) Empty() bool {
	return s.Len() == 0
}

func (s *Stack[t]) Len() int {
	return len(s.d)
}
