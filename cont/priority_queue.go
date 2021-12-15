package cont

// Build a new priorrity queue
func NewPriorityQueue[T any](cap int, less func(l, r T) bool) PriorityQueue[T] {
	var pq PriorityQueue[T]
	pq.d = NewHeap(cap, less)
	return pq
}

// Queue container
type PriorityQueue[T any] struct {
	d Heap[T]
}

// Push an element.
func (q *PriorityQueue[T]) Push(t T) {
	q.d.Push(t)
}

// Get an element from the queue. This function does not change state of the queue.
// If this function is being called on an empty queue, the panic will be raised.
func (q PriorityQueue[T]) Front() T {
	return q.d.GetMax()
}

// Pop an element.
func (q *PriorityQueue[T]) Pop() T {
	return q.d.Pop()
}

// Return true if the queue is empty.
func (q PriorityQueue[t]) Empty() bool {
	return q.d.Empty()
}

// Get the amount of elements in the queue.
func (q PriorityQueue[t]) Len() int {
	return q.d.Len()
}
