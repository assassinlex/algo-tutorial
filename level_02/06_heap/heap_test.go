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
	_ = heap.Push(100)
	_ = heap.Push(4)
	_ = heap.Push(20)
	heap.Println()
	v, _ := heap.Pop()
	fmt.Println("第一次弹出堆顶: ", v)
	heap.Println()

	v, _ = heap.Pop()
	fmt.Println("第二次弹出堆顶: ", v)
	heap.Println()
}

func TestSort(t *testing.T) {
	s := []int{500, 100, 233, 1, 44, 2, -100, 3, 10, 2}

	Sort(s)

	fmt.Println(s)
}
