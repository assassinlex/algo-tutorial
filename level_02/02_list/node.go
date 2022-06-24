package _2_list

// Node 节点类型
type Node interface {
	Prev() Node
	Next() Node
	SetPrev(n Node)
	SetNext(n Node)
	Value() int64
}

// ListNode 链表节点
type ListNode struct {
	value int64
	prev  Node
	next  Node
}

func (l *ListNode) Value() int64 {
	return l.value
}

func (l *ListNode) Prev() Node {
	return l.prev
}

func (l *ListNode) Next() Node {
	return l.next
}

func (l *ListNode) SetPrev(node Node) {
	l.prev = node
}

func (l *ListNode) SetNext(node Node) {
	l.next = node
}

func NewListNode(v int64) *ListNode {
	return &ListNode{
		value: v,
		prev:  nil,
		next:  nil,
	}
}

type TreeNode struct {
	value int64
	Left  *TreeNode
	Right *TreeNode
}

func (t *TreeNode) Prev() Node {
	panic("implement me")
}

func (t *TreeNode) Next() Node {
	panic("implement me")
}

func (t *TreeNode) SetPrev(n Node) {
	panic("implement me")
}

func (t *TreeNode) SetNext(n Node) {
	panic("implement me")
}

func (t *TreeNode) Value() int64 {
	return t.value
}

func NewTreeNode(v int64) *TreeNode {
	return &TreeNode{
		value: v,
		Left:  nil,
		Right: nil,
	}
}
