package queue

type Queue[T any] struct {
	storage []T
}

func (q *Queue[T]) Enqueue(val T) {
	q.storage = append(q.storage, val)

	return
}

func (q *Queue[T]) Dequeue() T {
	if q.IsEmpty() {
		panic("Queue is Empty")
	}

	val := q.storage[0]
	q.storage = q.storage[1:]

	return val
}

func (q *Queue[T]) Peek() T {
	if q.IsEmpty() {
		panic("Queue is Empty")
	}

	return q.storage[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.storage) == 0
}
