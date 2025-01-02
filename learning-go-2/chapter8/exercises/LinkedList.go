package main

type Node[T comparable] struct {
	value T
	next  *Node[T]
}
type LinkedList[T comparable] struct {
	head *Node[T]
}

func (l *LinkedList[T]) Add(value T) {
	newNode := Node[T]{value: value}
	if l.head == nil {
		l.head = &newNode
		return
	}
	current := l.head
	for current.next != nil {
		current = current.next
	}
	current.next = &newNode
}

func (l *LinkedList[T]) Insert(value T, position int) {
	newNode := Node[T]{value: value}
	if position == 0 {
		newNode.next = l.head
		l.head = &newNode
	}
	curr := l.head
	for i := 0; i < position-1 && curr != nil; i++ {
		curr = curr.next
	}
	if curr != nil {
		newNode.next = curr.next
		curr.next = &newNode
	}
}
func (l *LinkedList[T]) Index(value T) int {
	if l == nil {
		return -1
	}
	position := 0
	current := l.head
	for current != nil {
		if current.value == value {
			return position
		}
		current = l.head.next
		position++
	}

	return -1

}
func main() {

}
