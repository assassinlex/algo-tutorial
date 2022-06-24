package _2_list

import "sync/atomic"

// List 链表
type List interface {
	Head() Node
	Tail() Node
	Len() int64
}

// DoubleLinkList 双向链表
type DoubleLinkList struct {
	head   Node
	cur    Node
	tail   Node
	length int64
}

// Head 获取头节点
func (d *DoubleLinkList) Head() Node {
	return d.head
}

// Tail 获取尾节点
func (d *DoubleLinkList) Tail() Node {
	return d.tail
}

// Len 获取长度
func (d *DoubleLinkList) Len() int64 {
	return d.length
}

// PushFromHead  从头部添加元素
func (d *DoubleLinkList) PushFromHead(node Node) {
	if d.head == nil {
		d.head, d.cur, d.tail = node, node, node
		atomic.AddInt64(&d.length, 1)
		return
	}
	d.head.SetPrev(node)
	node.SetNext(d.head)
	d.head = node
	atomic.AddInt64(&d.length, 1)
}

// PushFromTail 从尾部添加元素
func (d *DoubleLinkList) PushFromTail(node Node) {
	if d.tail == nil {
		d.head, d.cur, d.tail = node, node, node
		atomic.AddInt64(&d.length, 1)
		return
	}
	d.tail.SetNext(node)
	node.SetPrev(d.tail)
	d.tail = node
	atomic.AddInt64(&d.length, 1)
}

func (d *DoubleLinkList) PopFromHead() Node {
	if d.Len() == 0 {
		return nil
	}

	head := d.Head()

	if d.Len() == 1 {
		d.head, d.tail = nil, nil
		atomic.AddInt64(&d.length, -1)
		return head
	}

	d.head = d.head.Next()
	d.head.SetPrev(nil)
	head.SetNext(nil)
	atomic.AddInt64(&d.length, -1)

	return head
}

func (d *DoubleLinkList) PopFromTail() Node {
	if d.Len() == 0 {
		return nil
	}

	tail := d.tail

	if d.Len() == 1 {
		d.head, d.tail = nil, nil
		atomic.AddInt64(&d.length, -1)
		return tail
	}

	d.tail = d.tail.Prev()
	d.tail.SetNext(nil)
	tail.SetPrev(nil)
	atomic.AddInt64(&d.length, -1)

	return tail
}

// NewDoubleLinkList 构造双向链表
func NewDoubleLinkList() *DoubleLinkList {
	return &DoubleLinkList{}

}
