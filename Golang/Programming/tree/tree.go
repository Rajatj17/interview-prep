package tree

import "reflect"

type MultiNodeTree[T any] struct {
	data     T
	children []*BinaryTree[T]
}

type BinaryTree[T any] struct {
	data      T
	leftNode  *BinaryTree[T]
	rightNode *BinaryTree[T]
}

type BinaryTreeCollector[T any] struct {
	root *BinaryTree[T]
}

func (root *BinaryTree[T]) insertNode(value T) {
	element := &BinaryTree[T]{data: value}

	if root == nil {
		root = element
		return
	}

	cmp := reflect.ValueOf(value).Interface().(T)
	rootData := reflect.ValueOf((*root).data).Interface().(T)

	if cmp < root.data {

	}
}

func main() {
	bC := BinaryTree[int]{}

	bC.insertNode(1)
}
