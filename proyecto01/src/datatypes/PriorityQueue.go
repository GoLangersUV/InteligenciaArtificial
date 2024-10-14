package datatypes

type PriorityQueue[T any] struct {
	elements []Element[T]
}

type Element[T any] struct {
	Value    T
	Priority int
}

func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{elements: make([]Element[T], 0)}
}

func (pq *PriorityQueue[T]) Push(element Element[T]) {
	pq.elements = append(pq.elements, element)
	pq.heapifyUp(len(pq.elements) - 1)
}

func (pq *PriorityQueue[T]) Pop() (T, bool) {
	if len(pq.elements) == 0 {
		var zero T
		return zero, false
	}
	root := pq.elements[0]
	pq.elements[0] = pq.elements[len(pq.elements)-1]
	pq.elements = pq.elements[:len(pq.elements)-1]
	pq.heapifyDown(0)
	return root.Value, true
}

func (pq *PriorityQueue[T]) heapifyUp(index int) {
	for index > 0 {
		parentIndex := (index - 1) / 2
		if pq.elements[parentIndex].Priority >= pq.elements[index].Priority {
			break
		}
		pq.elements[parentIndex], pq.elements[index] = pq.elements[index], pq.elements[parentIndex]
		index = parentIndex
	}
}

func (pq *PriorityQueue[T]) heapifyDown(index int) {
	for {
		leftChildIndex := 2*index + 1
		rightChildIndex := 2*index + 2
		largest := index
		if leftChildIndex < len(pq.elements) && pq.elements[leftChildIndex].Priority > pq.elements[largest].Priority {
			largest = leftChildIndex
		}
		if rightChildIndex < len(pq.elements) && pq.elements[rightChildIndex].Priority > pq.elements[largest].Priority {
			largest = rightChildIndex
		}
		if largest == index {
			break
		}
		pq.elements[largest], pq.elements[index] = pq.elements[index], pq.elements[largest]
		index = largest
	}
}

func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.elements) == 0
}
