package _2_list

/**
 * 双向链表
 */

// Node 节点
type Node struct {
	Value int64
	Prev *Node
	Next *Node
}

// DoubleLinkList 双向链表
type DoubleLinkList struct {
	head *Node
	tail *Node
	length int64
}

func NewNode(v int64) *Node {
	return &Node{
		Value: v,
	}
}

