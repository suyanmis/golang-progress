package main

import (
	"cmp"
	"fmt"
)

type OrderableFunc[T any] func(t1, t2 T) int

type Tree[T any] struct {
	f     OrderableFunc[T]
	root  *Node[T]
	count int
}

type Node[T any] struct {
	val         T
	left, right *Node[T]
}

func NewTree[T any](f OrderableFunc[T]) *Tree[T] {
	return &Tree[T]{
		f: f,
	}
}
func (t *Tree[T]) Add(v T) {
	t.root = t.root.Add(t.f, v)
	t.count++
}

func (t *Tree[T]) Contains(v T) bool {
	return t.root.Contains(t.f, v)
}

func (n *Node[T]) Add(f OrderableFunc[T], v T) *Node[T] {
	if n == nil {
		return &Node[T]{val: v}
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left = n.left.Add(f, v)
	default:
		n.right = n.right.Add(f, v)
	}
	return n
}

// if pointer is not used, nil will panic
func (n *Node[T]) Contains(f OrderableFunc[T], v T) bool {
	if n == nil {
		return false
	}
	switch r := f(v, n.val); {
	case r <= -1:
		n.left.Contains(f, v)
	case r >= 1:
		n.right.Contains(f, v)
	}
	return true

}

func (n *Node[T]) Traversal() []T {
	sum := make([]T, 0)

	if n == nil {
		return sum
	}
	if n.left == nil && n.right == nil {
		sum = append(sum, n.val)
		return sum
	}

	if n.left != nil {
		sum = append(sum, n.left.Traversal()...)
	}
	sum = append(sum, n.val)
	if n.right != nil {
		sum = append(sum, n.right.Traversal()...)
	}
	return sum
}
func (t *Tree[T]) Traversal() []T {
	if t.root == nil {
		return nil
	}
	return t.root.Traversal()

}
func main() {

	t1 := NewTree(cmp.Compare[int])
	t1.Add(15)
	t1.Add(10)
	t1.Add(30)
	t1.Add(15)
	fmt.Println(t1.Contains(15))
	fmt.Println(t1.Contains(40))

	fmt.Printf("t1.Traversal(): %v\n", t1.Traversal())

}
