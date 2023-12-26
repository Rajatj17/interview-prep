package priorityqueue

type Heap[T comparable] struct {
	arr []T
}

func (h *Heap[T]) Length() int {
	return len(h.arr)
}

func (h *Heap[T]) Swap() {

}

func (h *Heap[T]) Less(i, j int) bool {
	return true
}
