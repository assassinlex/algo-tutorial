package _5_stack

import (
	list "algo/level_02/02_list"
)

// Stack 栈
type Stack struct {
	list.DoubleLinkList
}

// Push 入栈
func (s *Stack) Push(node list.Node) {
	s.PushFromTail(node)
}

// Pop 出栈
func (s *Stack) Pop() list.Node {
	return s.PopFromTail()
}

// Size 栈大小
func (s *Stack) Size() int64 {
	return s.Len()
}

// NewStack 构造栈
func NewStack() *Stack {
	return &Stack{}
}
