package cont

// Build a new heap
func NewHeap[T any](cap int, less func(l, r T) bool) Heap[T] {
	var h Heap[T]
	h.less = less
	h.d = make([]T, 0, cap)
	return h
}

// Heap container
type Heap[T any] struct {
	d    []T
	less func(l, r T) bool
}

func (h Heap[T]) getLChild(i int) int {
	r := 2*i + 1
	if r >= len(h.d) {
		return -1
	}
	return r
}

func (h Heap[T]) getParent(i int) int {
	if i == 0 {
		return 0
	}
	return (i - 1) / 2
}

func (h *Heap[T]) siftUp(i int) {
	for {
		p := h.getParent(i)

		if h.less(h.d[p], h.d[i]) {
			// swap the parent with current child
			h.d[p], h.d[i] = h.d[i], h.d[p]
		} else {
			break
		}

		if p == 0 {
			break
		}

		i = p
	}
}

func (h *Heap[T]) siftDown(i int) {
	for {
		c := h.getLChild(i)

		if c == -1 {
			break
		}

		if c+1 < len(h.d) && h.less(h.d[c], h.d[c+1]) {
			c++
		}

		if !h.less(h.d[i], h.d[c]) {
			break
		}

		h.d[i], h.d[c] = h.d[c], h.d[i]
		i = c
	}
}

// Get max element. The heap must not be empty.
func (h Heap[T]) GetMax() T {
	return h.d[0]
}

// Remove max element and reurn it
func (h *Heap[T]) Pop() (r T) {
	l := len(h.d)
	if l == 0 {
		return
	}

	r = h.d[0]
	h.d[0] = h.d[l-1]
	h.d = h.d[:l-1]

	h.siftDown(0)
	return
}

// Push an element
func (h *Heap[T]) Push(v T) {
	h.d = append(h.d, v)
	h.siftUp(len(h.d) - 1)
}

// Get lenght
func (h Heap[T]) Len() int {
	return len(h.d)
}

// Check if the container is empty
func (h Heap[T]) Empty() bool {
	return h.Len() == 0
}
