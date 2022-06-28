package _2_list

type Node interface {
	Prev() Node
	SetPrev(Node)
	Next() Node
	SetNext(Node)
	Left() Node
	SetLeft(Node)
	Right() Node
	SetRight(Node)
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

func (l *ListNode) Left() Node {
	//TODO implement me
	return nil
}

func (l *ListNode) Right() Node {
	//TODO implement me
	return nil
}

func (l *ListNode) SetLeft(node Node) {
	//TODO implement me
}

func (l *ListNode) SetRight(node Node) {
	//TODO implement me
}

// NewListNode 初始化节点
func NewListNode(v int64) *ListNode {
	return &ListNode{
		value: v,
		prev:  nil,
		next:  nil,
	}
}

// TreeNode 树型节点
type TreeNode struct {
	*ListNode
	left  Node
	right Node
}

func (t *TreeNode) SetLeft(node Node) {
	t.left = node
}

func (t *TreeNode) SetRight(node Node) {
	t.right = node
}

func (t *TreeNode) Left() Node {
	return t.left
}

func (t *TreeNode) Right() Node {
	return t.right
}

// NewTreeNode 构造树型节点
func NewTreeNode(v int64) *TreeNode {
	return &TreeNode{
		ListNode: &ListNode{value: v},
	}
}
