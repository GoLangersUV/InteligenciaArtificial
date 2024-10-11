package datatypes

type Queue[T any] struct {
	elements []T
}

func (q *Queue[T]) Enqueue(v T) {
	q.elements = append(q.elements, v)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, true
	}
	v := q.elements[0]
	q.elements = q.elements[1:]
	return v, false
}

func (q *Queue[T]) Len() int {
	return len(q.elements)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.elements) == 0
}
