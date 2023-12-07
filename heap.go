package ticktock

func getLeftChildIndex(idx int) int {
	return idx*2 + 1
}

func getRightChildIndex(idx int) int {
	return idx*2 + 2
}

func getParentIndex(idx int) int {
	return (idx - 1) / 2
}

type Heap[T any] struct {
	f     OrderableFunc[T]
	array []T
}

// NewHeap creates a new Heap
func NewHeap[T any](of OrderableFunc[T]) *Heap[T] {
	return &Heap[T]{f: of}
}

func (h *Heap[T]) Len() int {
	return len(h.array)
}

func (h *Heap[T]) getLeftChild(idx int) T  { return h.array[getLeftChildIndex(idx)] }
func (h *Heap[T]) getRightChild(idx int) T { return h.array[getRightChildIndex(idx)] }
func (h *Heap[T]) getParent(idx int) T     { return h.array[getParentIndex(idx)] }
func (h *Heap[T]) swap(i, j int)           { h.array[i], h.array[j] = h.array[j], h.array[i] }

// IsEmpty returns true if the heap is empty
func (h *Heap[T]) IsEmpty() bool {
	return len(h.array) == 0
}

// Insert inserts a value into the heap
// and heapifies it
func (h *Heap[T]) Insert(value T) {
	h.array = append(h.array, value)
	h.heapUp()
}

// Remove removes the root element from the heap
// and heapifies it
func (h *Heap[T]) Remove() T {
	var root T
	if len(h.array) == 0 {
		return root
	}

	root = h.array[0]
	h.array[0] = h.array[len(h.array)-1]
	h.array = h.array[:len(h.array)-1]
	h.heapDown()

	return root
}

// Peek returns the root element of the heap
func (h *Heap[T]) Peek() T {
	return h.array[0]
}

// heapDown moves the root element down the heap
func (h *Heap[T]) heapDown() {
	idx := 0
	for {
		minChild := getLeftChildIndex(idx)

		if minChild >= len(h.array) {
			return
		}

		if getRightChildIndex(idx) < len(h.array) && h.f(h.getLeftChild(idx), h.getRightChild(idx)) > 0 {
			minChild = getRightChildIndex(idx)
		}

		if h.f(h.array[idx], h.array[minChild]) < 0 {
			return
		}

		h.swap(idx, minChild)
		idx = minChild
	}
}

// heapUp moves an element up the heap
func (h *Heap[T]) heapUp() {
	idx := len(h.array) - 1

	for idx > 0 {
		if h.f(h.array[idx], h.getParent(idx)) > 0 {
			break
		}

		h.swap(idx, getParentIndex(idx))
		idx = getParentIndex(idx)
	}
}
