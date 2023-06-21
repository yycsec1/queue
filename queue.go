package queue

type Queue[T any] struct {
	items []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (queue *Queue[T]) Push(item T) {
	queue.items = append(queue.items, item)
}

func (queue *Queue[T]) Pop() (T, bool) {
	var returnValue T
	length := len(queue.items)
	if length > 0 {
		returnValue, queue.items = queue.items[0], queue.items[1:]
		return returnValue, true
	}
	return returnValue, false
}

func (queue *Queue[T]) Peek() (T, bool) {
	var returnValue T
	length := len(queue.items)
	if length > 0 {
		returnValue = queue.items[0]
		return returnValue, true
	}
	return returnValue, false
}

func (queue *Queue[T]) IsEmpty() bool {
	return len(queue.items) == 0
}

func (queue *Queue[T]) Len() int {
	return len(queue.items)
}
