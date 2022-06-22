// 单向链表、双向链表

package level_01

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
)

func TestOneWayReserve(t *testing.T) {
	n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	n3 := &Node{value: 3}
	n4 := &Node{value: 4}
	n1.next = n2
	n2.next = n3
	n3.next = n4

	head := n1

	for head != nil {
		log.Printf("%v\n", head.value)
		head = head.next
	}

	head = oneWayLinkedListReverse(n1)

	for head != nil {
		log.Printf("%v\n", head.value)
		head = head.next
	}
}

func TestDoubleLinkedListReverse(t *testing.T)  {
	n1 := &Node{value: 1}
	n2 := &Node{value: 2}
	n3 := &Node{value: 3}
	n4 := &Node{value: 4}
	n1.next = n2
	n2.prev = n1
	n2.next = n3
	n3.prev = n2
	n3.next = n4
	n4.prev = n3

	head := n1

	for head != nil {
		log.Printf("value: %d\tself.address: %v\tprev.address: %v\tnext.address: %v\n",
			head.value, head, head.prev, head.next)
		head = head.next
	}

	head = doubleLinkedListReverse(n1)

	for head != nil {
		log.Printf("value: %d\tself.address: %v\tprev.address: %v\tnext.address: %v\n",
			head.value, head, head.prev, head.next)
		head = head.next
	}
}

func TestListNodeSum(t *testing.T)  {
	// 随机链表长度值
	long, short := rand.Intn(10), rand.Intn(10)
	if short > long {
		long, short = short, long
	}
	// 根据长度生成随机值链表
	l := generateListNode(long)
	s := generateListNode(short)
	// 复制长链表，因为测试结果会改变原有链表
	tl := copyListNode(l)
	ts := s
	// 测试结果
	res := twoListSum(l, s)
	// 输出原有长、短，结果链表
	for tl != nil {
		fmt.Printf("%d\t", tl.val)
		tl = tl.next
	}
	println()
	for ts != nil {
		fmt.Printf("%d\t", ts.val)
		ts = ts.next
	}
	println()
	for res != nil {
		fmt.Printf("%d\t", res.val)
		res = res.next
	}
}

func TestSortedListMerge(t *testing.T) {
	data := [][]int {
		{1, 1, 5, 7, 9},
		{0, 7, 3, 6, 8},
	}
	list1 := generateSortedListNode(data[0])
	list2 := generateSortedListNode(data[1])
	tmpL1, tmpL2 := list1, list2
	for tmpL1 != nil {
		fmt.Printf("%d\t", tmpL1.val)
		tmpL1 = tmpL1.next
	}
	println()
	for tmpL2 != nil {
		fmt.Printf("%d\t", tmpL2.val)
		tmpL2 = tmpL2.next
	}
	println()
	res := twoSortedListMerge(list1, list2)
	for res != nil {
		fmt.Printf("%d\t", res.val)
		res = res.next
	}
}