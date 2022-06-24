package _5_stack

import (
	list "algo/level_02/02_list"
	"fmt"
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack([]*list.Node{
		list.NewNode(2),
		list.NewNode(4),
		list.NewNode(6),
		list.NewNode(8),
		list.NewNode(10),
		list.NewNode(12),
	})

	for s.Size() > 0 {
		fmt.Printf("弹出前栈大小: %d, 弹出当前栈顶元素: %d, 弹出后栈大小: %d\n", s.Size(), s.Pop().Value, s.Size())
	}
}
