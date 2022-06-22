// 单向链表、双向链表

package level_01

import (
	"math/rand"
	"sort"
)

type Node struct {
	value int
	prev *Node
	next *Node
}

type ListNode struct {
	val int
	next *ListNode
}

// 单向量表反转
// 初始链表
//    nil    ->    N1    ->    N2    ->    N3    ->    nil
// 循环条件 head 不能为空
//
// 1. 记录当前 head(N1) 的 prev(nil) 和 next(N2)，如上
//    next = head.next
//    nil    ->    N1    ->    N2    ->    N3    ->    nil
//    prev        head        next
// 2. 改变当前 head 的 next 指向，指向其 prev，此时 N1 & N2 的链接断开，N1.next 变为 nil
//    head.next = prev
//    nil    <-    N1          N2    ->    N3    ->    nil
//    prev        head        next
// 3. 移动 prev 到当前 head，移动当前 head 到 当前 next
//    prev = head
//    head = next
//    nil    <-    N1          N2    ->    N3    ->    nil
//                prev        head
//                            next
// 4. 开启下一轮循环(head != nil)
//    next = head.next
//    nil    <-    N1          N2    ->    N3    ->    nil
//                prev        head        next
//    head.next = prev
//    nil    <-    N1    <-    N2          N3    ->    nil
//                prev        head        next
//    prev = head
//    head = next
//    nil    <-    N1    <-    N2          N3    ->    nil
//                            prev        head
//                                        next
//    下一轮循环(head != nil)
//    next = head.next
//    nil    <-    N1    <-    N2          N3    ->    nil
//                            prev        head        next
//    head.next = prev
//    nil    <-    N1    <-    N2    <-    N3          nil
//                            prev        head        next
//    prev = head
//    head = next
//    nil    <-    N1    <-    N2    <-    N3    <-    nil
//                                        prev        head
//                                                    next
// 5. head == nil 跳出循环，返回 prev 节点由外部持有，程序结束


// oneWayLinkedListReverse 单向链表反转
func oneWayLinkedListReverse(head *Node) *Node {
	var prev, next *Node
	for head != nil {
		// 1. 锁住 head.next
		next = head.next
		// 2. head.next 反转 (删除 head 原有 next 指向，改为 prev 指向)
		head.next = prev
		// 3. 移动 head & prev 指向 (将 prev 指向当前 head，将当前 head 指向 next，开始下一个循环)
		prev = head
		head = next
	}

	return prev
}


// doubleLinkedListReverse 双向链表反转
func doubleLinkedListReverse(head *Node) *Node {
	var prev, next *Node
	for head != nil {
		// 1. 锁住 head.next
		next = head.next
		// 2. head.next & head.prev 反转
		head.next = prev
		head.prev = next
		// 移动 prev & head
		prev = head
		head = next
	}

	return prev
}


// 两个链表相加
// 给定两个链表的头结点 head1 & head2, 认为从左到右是某个数字的从低位到高位, 返回相加之后的链表
//  		4 -> 3 -> 6
//  		2 -> 5 -> 3
//  		6 -> 8 -> 9
// 解释:    634 + 252 = 986
//
// 解题思路:
//    1. 链表又长又短, 标记长短链表
//    2. 分别对位节点上的值相加可能产生进位，需要一个变量保存进位
//    3. 有三个阶段相加
//        ①: 长短链表都在相加
//        ②: 短链表结束, 长短链表与进位值相加
//        ③: 长链表结束, 但是还有进位值, 此时需要在长节点末尾添加一个节点, 值为进位值
func twoListSum(head1, head2 *ListNode) *ListNode {
	// 获取两条链表长度
	length1 := listLength(head1)
	length2 := listLength(head2)
	// 定义长短链表头节点
	longListNode := head1
	shortListNode := head2
	if length2 > length1 {
		longListNode, shortListNode = head2, head1
	}
	// 定义长短链表当前头节点
	curLongNode := longListNode
	curShortNode := shortListNode
	// 定义长链表的尾结点
	lastLongNode := curLongNode
	// 定义进位值
	carry := 0
	// 第一阶段
	for curShortNode != nil {
		// 当前所在节点值加和: 长 + 短 + 进位
		sum := curLongNode.val + curShortNode.val + carry
		// 更新当前节点值 & 进位值，如 9 + 2 = 11，当前值为 1，进位值为 1
		curLongNode.val = sum % 10
		carry = sum / 10
		// 更新长链表尾结点
		lastLongNode = curLongNode
		// 更新当前节点位置
		curLongNode = curLongNode.next
		curShortNode = curShortNode.next
	}
	// 第二阶段
	for curLongNode != nil {
		// 长节点 & 进位值的加和
		sum := curLongNode.val + carry
		// 更新当前节点值 & 进位值，如 9 + 1 = 10，当前值为 0，进位值为 1
		curLongNode.val = sum % 10
		carry = sum / 10
		// 更新长链表尾结点
		lastLongNode = curLongNode
		// 更新当前节点位置
		curLongNode = curLongNode.next
	}
	// 第三阶段
	if carry != 0 {
		lastLongNode.next = &ListNode{val: carry}
	}
	// 将更新后的长链表作为结果返回
	return longListNode
}

// twoSortedListMerge 合并两个有序链表
// 1 -> 3 -> 5 -> 7 -> 9    0 -> 2 -> 4 -> 6 -> 8    =>    0 -> 1 -> 2 -> ... -> 8 -> 9
func twoSortedListMerge(head1, head2 *ListNode) *ListNode {
	// 若有链表为空, 返回另一个链表
	if head1 == nil || head2 == nil {
		if head1 == nil {
			return head2
		} else {
			return head1
		}
	}
	// 先决定 prev 节点 & 头节点: 两个链表的头节点的大小关系, 较小的头作为返回值
	prev := head1
	if head1.val > head2.val {
		prev = head2
	}
	head := prev
	// 标记两个当前节点
	cur1 := head.next
	cur2 := head1
	if head == head1 {
		cur2 = head2
	}
	// 比较当前节点值的大小，谁小作为新链表的下一个节点, 更新当前节点
	// 记录 prev 节点，用来指向新链表的下一个节点
	for cur1 != nil && cur2 != nil {
		if cur1.val > cur2.val {
			prev.next = cur2
			cur2 = cur2.next
		} else {
			prev.next = cur1
			cur1 = cur1.next
		}
		prev = prev.next
	}
	// 当短链表遍历完后，将 prev 节点指向长链表包含当前节点在内的后续所有节点
	if cur1 != nil || cur2 != nil {
		if cur1 != nil {
			prev.next = cur1
		} else {
			prev.next = cur2
		}
	}
	return head
}



// listLength 链表长度
func listLength(head *ListNode) int {
	length := 0
	for head != nil {
		length++
		head = head.next
	}
	return length
}

// generateListNode 生成随机值链表
func generateListNode(length int) *ListNode {
	val := rand.Intn(10)
	head := &ListNode{val: val}
	curHead := head
	for i := 0; i < length - 1; i++ {
		val := rand.Intn(10)
		curHead.next = &ListNode{val: val}
		curHead = curHead.next
	}

	return head
}

// generateListNode 生成随机值链表
func generateSortedListNode(s []int) *ListNode {
	sort.Ints(s)
	length := len(s)
	if length == 0 {
		return &ListNode{val: 0}
	}
	head := &ListNode{val: s[0]}
	curHead := head
	for i := 1; i < length; i++ {
		curHead.next = &ListNode{val: s[i]}
		curHead = curHead.next
	}
	return head
}

// copyListNode 复制链表
func copyListNode(head *ListNode) *ListNode {
	tmpHead := &ListNode{val: head.val}
	newHead := tmpHead
	head = head.next
	for head != nil {
		tmpHead.next = &ListNode{val: head.val}
		head = head.next
		tmpHead = tmpHead.next
	}
	return newHead
}
