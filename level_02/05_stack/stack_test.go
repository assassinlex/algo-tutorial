package _5_stack

import (
	list "algo/level_02/02_list"
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack()

	for i := 1; i < 10; i += 2 {
		s.Push(list.NewListNode(int64(i)))
	}

	for s.Size() > 0 {
		fmt.Printf("弹出前栈大小: %d, 弹出当前栈顶元素: %d, 弹出后栈大小: %d\n", s.Size(), s.Pop().Value(), s.Size())
	}
}
