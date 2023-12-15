package stack

import linkedlist "programming/linked-list"

type LLStack[T any] struct {
	list *linkedlist.List[T]
}
