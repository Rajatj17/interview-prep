package stack

func NewMinMax[T any](compare func(i, j T) bool, values ...T) *MinMaxStack[T] {
	stack := MinMaxStack[T]{
		valuesStack: New[T](),
		minMaxStack: New[[]T](),
		compare:     compare,
	}

	return &stack
}

type MinMaxStack[T any] struct {
	valuesStack *SliceStack[T]
	minMaxStack *SliceStack[[]T]
	compare     func(i, j T) bool
}
