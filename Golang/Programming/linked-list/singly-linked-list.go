package linkedlist

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	val  interface{}
	next *Node[T]
}

type List[T any] struct {
	count int
	tail  *Node[T]
	head  *Node[T]
}

func New[T any](values ...T) *List[T] {
	list := List[T]{}
	list.Add(values...)

	return &list
}

func (list *List[T]) isEmpty() bool {
	return list.count <= 0
}

func (list *List[T]) Size() int {
	return list.count
}

func (list *List[T]) Add(values ...T) {
	// for _, values := range values {}
}

// Normal append, Assuming no tail ptr or count variable
func (list *List[T]) Append(value T) {
	element := &Node[T]{val: value}

	if list.head == nil {
		list.head = element

		return
	}

	head := list.head
	for head.next != nil {
		head = head.next
	}

	head.next = element

	return
}

func (list *List[T]) InsertAt(index int, value T) error {
	if index < 0 {
		return errors.New("Ye Kya backhodi h")
	}

	if index > list.count {
		return errors.New("Out of Bound Index")
	}

	// Create New Node
	newNode := &Node[T]{val: value}
	if list.isEmpty() {
		list.head = newNode
		list.tail = newNode
		list.count++

		// Make newnode nil to be garbage collected & avoid un-intential write off
		newNode = nil

		return nil
	}

	current, _ := list.getNode(index - 1)

	newNode.next = current.next
	current.next = newNode
	list.count++

	// Make newNode nil to be garbage collected & avoid un-intential write off
	current = nil
	newNode = nil

	return nil
}

func (list *List[T]) getNode(index int) (res *Node[T], err error) {
	current := list.head
	for i := 0; i < index; i++ {
		current = current.next
	}

	return current, nil
}

func (list *List[T]) RemoveAt(index int) (res T, err error) {
	if index <= 0 {
		return res, errors.New("Invalid")
	}

	if index == 0 {
		current := list.head
		list.head = current.next

		res = current.val.(T)

		current.next = nil
		current = nil
		list.count--

		return res, nil
	}

	previousElement, _ := list.getNode(index - 1)
	current := previousElement.next

	res = current.val.(T)

	previousElement.next = current.next
	current = nil
	list.count--

	return res, nil
}

func (list *List[T]) GetValues() []T {
	values := []T{}

	current := list.head
	if current == nil {
		return values
	}

	for current.next != nil {
		values = append(values, current.val.(T))
		current = current.next
	}

	values = append(values, current.val.(T))

	return values
}

func main() {
	var nodeValue string
	linkList := &List[string]{}

	for {
		fmt.Scanln(&nodeValue)
		if nodeValue == "-1" {
			break
		}

		linkList.Append(nodeValue)
	}

	res := linkList.GetValues()
	fmt.Println(res)
}
