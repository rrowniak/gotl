package gotl

// Queue container
type Queue[T any] struct {
	d     []T
	front int
}

// Increase the capacity of the queue. If the queue is not empty, the elements will be copied to the new location.
func (q *Queue[T]) Reserve(size int) {
	if q.Empty() {
		q.d = make([]T, 0, size)
	} else {
		new := make([]T, len(q.d), size)
		copy(new, q.d)
		q.d = new
	}
}

func (q *Queue[T]) shrink() {
	if q.front == 0 {
		// nothing to fo
		return
	}
	if !q.Empty() {
		for i := 0; i < q.Len(); i++ {
			q.d[i] = q.d[i+q.front]
		}
	}
	q.d = q.d[0:q.Len()]
	q.front = 0
}

// Push an element.
func (q *Queue[T]) Push(t T) {
	q.d = append(q.d, t)
}

// Get an element from the queue. This function does not change state of the queue.
// If this function is being called on an empty queue, the panic will be raised.
func (q Queue[T]) Front() T {
	return q.d[q.front]
}

// Pop an element.
func (q *Queue[T]) Pop() {
	if q.Len() == 0 {
		// nothing to do
		return
	}

	q.front++

	// time to shrink?
	if q.front > (cap(q.d) >> 1) {
		q.shrink()
	}
}

// Return true if the queue is empty.
func (q Queue[t]) Empty() bool {
	return q.Len() == 0
}

// Get the amount of elements in the queue.
func (q Queue[t]) Len() int {
	return len(q.d) - q.front
}
