package _6_heap

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	heap := NewMaxHeap(10)
	_ = heap.Push(10)
	_ = heap.Push(1)
	_ = heap.Push(2)
	_ = heap.Push(5)
	_ = heap.Push(3)
	_ = heap.Push(9)
	heap.Println()
	v, _ := heap.Pop()
	fmt.Println(v)
	heap.Println()
}
