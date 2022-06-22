package _4_queue

import (
	list "algo/level_02/02_list"
	"sync/atomic"
)

type queue interface {
	Head() int64
	Tail() int64
	Push(v int64)
	Pop() int64
	Len() int64
	Empty() bool
}

type Queue struct {
	head *list.Node
	tail *list.Node
	length int64
}

func (q *Queue) Head() int64 {
	return q.head.Value
}

func (q *Queue) Tail() int64 {
	return q.tail.Value
}

func (q *Queue) Push(v int64) {
	node := list.NewNode(v)
	if q.head == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.Next = node
		node.Prev = q.tail
		q.tail = node
	}
	atomic.AddInt64(&(q.length), 1)
}

func (q *Queue) Pop() int64 {
	var ret int64

	if q.head == nil {
		// 队列为空
		return 0
	} else if q.head == q.tail {
		// 队列只有一个元素
		ret = q.head.Value
		q.head = nil
		q.tail = nil
	} else {
		head := q.head
		ret = head.Value
		q.head = head.Next
		q.head.Prev = nil
		head.Next = nil
	}

	atomic.AddInt64(&(q.length), -1)

	return ret
}

func (q *Queue) Len() int64 {
	return q.length
}

func (q *Queue) Empty() bool {
	return q.length == 0
}

func NewQueue(s []int64) *Queue {
	q := &Queue{}

	for _, v := range s {
		q.Push(v)
	}

	return q
}

