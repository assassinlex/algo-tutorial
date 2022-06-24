package _2_list

import (
	"testing"
)

func TestNewDoubleLinkList(t *testing.T) {
	list := NewDoubleLinkList()

	for i := 1; i < 10; i += 2 {
		list.PushFromTail(NewListNode(int64(i)))
	}

	head := list.Head()

	t.Log("从头部开始遍历")
	for head != nil {
		t.Logf("%d\t", head.Value())
		head = head.Next()
	}

	t.Log("从尾部开始遍历")
	tail := list.Tail()

	for tail != nil {
		t.Logf("%d\t", tail.Value())
		tail = tail.Prev()
	}
}
