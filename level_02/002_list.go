package level_02

import "math"

/**
 * - 链表结构
 * - 栈
 * - 队列
 * - 递归行为
 * - 哈希表
 */

type Node struct {
	value int
	prev *Node
	next *Node
}

// 生成单向链表
func generateSingleList(s []int) *Node {
	length := len(s)
	if length <= 0 {
		return nil
	}

	head := &Node{ value: s[0], prev: nil, next: nil}
	curHead := head

	for i := 1; i < length; i++ {
		node := &Node{ value: s[i], prev: nil, next: nil}
		curHead.next = node
		curHead = node
	}

	return head
}

// 单向链表反转
func singleListReverse(head *Node) *Node {
	var prev, next *Node

	for head != nil {
		// 定义 next
		next = head.next
		// 定义 head.next
		head.next = prev
		// 定义 prev
		prev = head
		// 移动 head
		head = next
	}

	return prev
}

// todo:: 双向链表反转

// Queue 队列 [实现: 双向链表]
type Queue struct {
	head *Node
	tail *Node
}

// NewQueue 获取队列实例
func NewQueue(n *Node) *Queue {
	return &Queue{ head: n, tail: n }
}

// 入队
func (q *Queue) enqueue(n *Node) {
	if q.isEmpty() {
		q.head = n
		q.tail = n
		return
	}
	curTail := q.tail
	curTail.next = n
	n.prev = curTail
	q.tail = n
}

// 出队
func (q *Queue) dequeue() *Node {
	if q.isEmpty() {
		return nil
	}
	// 临时记录头结点, 找到次头结点, 标记为新的头结点, 更新指向, 返回临时头结点
	curHead := q.head
	// 最后一个元素
	if q.head == q.tail {
		q.head, q.tail = nil, nil
	} else {
		q.head = curHead.next
		q.head.prev = nil

	}
	return curHead
}

// 队列是否为空
func (q *Queue) isEmpty() bool {
	return q.head == nil
}

// Stack 栈 [链表竖线]
type Stack struct {
	head *Node
}

// NewStack 获取栈实例
func NewStack(n *Node) *Stack {
	return &Stack{head: n}
}

// 入栈
func (s *Stack) push(n *Node) {
	if s.isEmpty() {
		s.head = n
	} else {
		curHead := s.head
		curHead.prev = n
		n.next = curHead
		s.head = n
	}
}

// 出栈
func (s *Stack) pop() *Node {
	if s.isEmpty() {
		return nil
	}
	curHead := s.head
	s.head = curHead.next
	if s.head != nil {
		s.head.prev = nil
	}
	return curHead
}

// 队列是否为空
func (s *Stack) isEmpty() bool {
	return s.head == nil
}

// todo:: 数组实现队列、栈 [队列: 环形数组]


// todo:: 栈最小值 [额外维护一个最小栈]


// 递归求最大值
func getMax(s []int, left, right int) int {
	if left == right {
		return s[left]
	}
	// mid := left + ((right >> 2) + 1)
	mid := left + ((right - left) >> 1)
	maxLeft := getMax(s, left, mid)
	maxRight := getMax(s, mid + 1, right)
	return int(math.Max(float64(maxLeft), float64(maxRight)))
}


// todo:: 哈希表