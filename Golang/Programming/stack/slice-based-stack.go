package stack

import "fmt"

type SliceStack[T any] struct {
	array []T
}

func New[T any](values ...T) *SliceStack[T] {
	stack := SliceStack[T]{make([]T, 0, len(values))}
	// stack.Push(values...)
	return &stack
}

func (s *SliceStack[T]) Push(val T) {
	s.array = append(s.array, val)
}

func (s *SliceStack[T]) Pop() (val T) {
	var lastIndex int = len(s.array) - 1
	val = s.array[lastIndex]

	s.array = s.array[0:lastIndex]
	return
}

func (s *SliceStack[T]) Size() int {
	return len(s.array)
}

func (s *SliceStack[T]) Peek() (val T) {
	return s.array[len(s.array)-1]
}

func main() {
	stack := &SliceStack[string]{}

	var stackValue string
	for {
		fmt.Scanln(&stackValue)
		if stackValue == "exit" {
			break
		}
		stack.Push(stackValue)
	}
}
