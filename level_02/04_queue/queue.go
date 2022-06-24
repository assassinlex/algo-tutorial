package _4_queue

import (
	list "algo/level_02/02_list"
)

// Queue 队列
type Queue struct {
	list.DoubleLinkList
}

// Enqueue 入队
func (q *Queue) Enqueue(node list.Node) {
	q.PushFromTail(node)
}

// Dequeue 出队
func (q *Queue) Dequeue() list.Node {
	return q.PopFromHead()
}

// NewQueue 构造队列
func NewQueue() *Queue {
	return &Queue{}
}
