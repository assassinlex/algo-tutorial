package day_02

import "errors"

/*
	用链表实现栈 & 队列
		双端队列: 用双向链表实现, 包含头尾两个节点
		栈:   从头入, 从头出 [后进先出]
        队列: 从头入, 从尾出 [先进先出]
*/

func NewDoubleEndsQueue() *doubleEndsQueue {
	return &doubleEndsQueue{}
}

// 双端队列
type doubleEndsQueue struct {
	head *doubleNode
	tail *doubleNode
}

// addFromHead 从链表头部添加元素
func (q *doubleEndsQueue) addFromHead(v int) {
	_node := &doubleNode{v: v}
	if q.head == nil {
		q.head = _node
		q.tail = _node
	} else {
		_node.next = q.head
		q.head.prev = _node
		q.head = _node
	}
}

// addFromBottom 从链表尾部添加
func (q *doubleEndsQueue) addFromBottom(v int) {
	_node := &doubleNode{v: v}
	if q.head == nil {
		q.head = _node
		q.tail = _node
	} else {
		_node.prev = q.tail
		q.tail.next = _node
		q.tail = _node
	}
}

// popFromHead 从头部弹出
func (q *doubleEndsQueue) popFromHead() *doubleNode {
	if q.head == nil {
		return nil
	}
	_node := q.head
	if q.head == q.tail {
		q.head, q.tail = nil, nil
	} else {
		q.head = q.head.next
		_node.next, _node.prev = nil, nil
	}
	return _node
}

// popFromBottom 从尾部弹出
func (q *doubleEndsQueue) popFromBottom() *doubleNode {
	if q.tail == nil {
		return nil
	}
	_node := q.tail
	if q.head == q.tail {
		q.head, q.tail = nil, nil
	} else {
		q.tail = q.tail.prev
		q.tail.next, _node.prev = nil, nil
	}
	return _node
}

// isEmpty 数据结构是否为空
func (q *doubleEndsQueue) isEmpty() bool {
	return q.head == nil
}

func NewStackByList() *stackByList {
	return &stackByList{queue: NewDoubleEndsQueue()}
}

type stackByList struct {
	queue *doubleEndsQueue
}

// Push 入栈
func (s *stackByList) Push(v int) {
	s.queue.addFromHead(v)
}

// Pop 出栈
func (s *stackByList) Pop() (int, error) {
	_node := s.queue.popFromHead()
	if _node == nil {
		return 0, errors.New("暂无数据")
	}
	return _node.v, nil
}

// Peek 获取栈顶, 但不出栈
func (s stackByList) Peek() (int, error) {
	_node := s.queue.head
	if _node == nil {
		return 0, errors.New("暂无数据")
	}
	return _node.v, nil
}

// IsEmpty 栈是否为空
func (s *stackByList) IsEmpty() bool {
	return s.queue.isEmpty()
}

func NewQueueByList() *queueByList {
	return &queueByList{
		queue: NewDoubleEndsQueue(),
		size:  0,
	}
}

type queueByList struct {
	queue *doubleEndsQueue
	size  int
}

// Push 入队
func (q *queueByList) Push(v int) {
	q.queue.addFromHead(v)
	q.size++
}

// Poll 出队
func (q *queueByList) Poll() (int, error) {
	_node := q.queue.popFromBottom()
	if _node == nil {
		return 0, errors.New("暂无数据")
	}
	q.size--
	return _node.v, nil
}

func (q *queueByList) Size() int {
	return q.size
}
