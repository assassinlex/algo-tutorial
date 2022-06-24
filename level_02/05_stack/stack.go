package _5_stack

import (
	list "algo/level_02/02_list"
	"sync/atomic"
)

// Stack 栈
type Stack struct {
	top    *list.Node
	bottom *list.Node
	size   int64
}

func (s *Stack) Push(node *list.Node) {
	if s.top == nil {
		s.top = node
		s.bottom = node
	} else {
		s.top.Next = node
		node.Prev = s.top
		s.top = node
	}

	atomic.AddInt64(&s.size, 1)
}

func (s *Stack) Pop() *list.Node {
	if s.top == nil {
		return s.top
	}
	ret := s.top
	if s.top == s.bottom {
		s.top = nil
		s.bottom = nil
	} else {
		top := s.top
		s.top = top.Prev
		top.Prev = nil
		s.top.Next = nil
	}

	atomic.AddInt64(&s.size, -1)
	return ret
}

func (s *Stack) Size() int64 {
	return s.size
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func NewStack(n []*list.Node) *Stack {
	s := &Stack{}

	for _, node := range n {
		s.Push(node)
	}

	return s
}
