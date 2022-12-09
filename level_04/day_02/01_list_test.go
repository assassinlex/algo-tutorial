package day_02

import "testing"

func TestReservedLinkedList(t *testing.T) {
	s := []int{
		1, 2, 3, 4, 5,
	}
	head := NewSingleLinkedList(s)
	_head := head
	traverseSingle(_head)
	reserved := reserveLinkedList(head)
	traverseSingle(reserved)
}

func TestReservedDoubleNode(t *testing.T) {
	s := []int{
		1, 2, 3, 4, 5,
	}
	head := NewDoubleLinkedList(s)
	_head := head
	traverseDouble(_head)
	reserved := reserveDoubleLinkedList(head)
	traverseDouble(reserved)
}

func TestDeleteValue(t *testing.T) {
	s := []int{
		1, 2, 2, 1, 3, 4, 1, 5,
	}
	head1 := NewSingleLinkedList(s)
	traverseSingle(head1)
	res1 := deleteValue(head1, 1)
	traverseSingle(res1)

	head2 := NewSingleLinkedList(s)
	res2 := deleteValue(head2, 2)
	traverseSingle(res2)

	head3 := NewSingleLinkedList(s)
	res3 := deleteValue(head3, 3)
	traverseSingle(res3)
}
