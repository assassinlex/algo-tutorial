package day_02

import "fmt"

/*
	链表: 一种线性表, 但是不按线性顺序存储数据, 而是在每一个数据节点中存储下一个节点地址指针.
		 链表在插入节点的时间复杂度为O(1), 查找节点时间复杂度为O(N).

		1. 单向链表: 每个节点包含值 & 下一个节点的指针, 只能由当前节点访问下一个节点. 尾结点的指针指向空(null)
		2. 双向链表: 每个节点包含值 & 前后两个节点的指针, 当前节点可以访问前后两个节点, 头结点的头指针指向空, 尾结点的尾指针指向空
		3. 循环链表: todo 待拓展...
		4. 块状链表: todo 待拓展...
*/

// 单向链表节点
type node struct {
	v    int
	next *node
}

// 双向链表节点
type doubleNode struct {
	v    int
	prev *doubleNode
	next *doubleNode
}

func NewSingleLinkedList(s []int) *node {
	if len(s) == 0 {
		return nil
	}
	head := &node{v: s[0]}
	_head := head
	for i := 1; i < len(s); i++ {
		_node := &node{v: s[i]}
		_head.next = _node
		_head = _node
	}
	return head
}

func NewDoubleLinkedList(s []int) *doubleNode {
	if len(s) == 0 {
		return nil
	}
	head := &doubleNode{v: s[0]}
	_head := head
	for i := 1; i < len(s); i++ {
		_node := &doubleNode{v: s[i], prev: _head}
		_head.next = _node
		_head = _node
	}
	return head
}

func traverseSingle(head *node) {
	for head != nil {
		fmt.Print(head.v, "\t")
		head = head.next
	}
	fmt.Println()
}

func traverseDouble(head *doubleNode) {
	for head != nil {
		fmt.Print(head.v, "\t")
		head = head.next
	}
	fmt.Println()
}

/*
 * 单向链表反转[ -> 代表 next 指针]
 *     null  ->  a  ->  b  ->  c -> null <入参>
 *     null  ->  c  ->  b  ->  a -> null <出参>
 * 解题步骤:
 *   1. 初始化 prev, next 两个节点, 分别代表入参链表当前节点的前一个节点 & 后一个节点
 *   2. 从头部节点开始遍历整个链表, 直至当前节点为 null, 代表整个链表处理完毕
 *   3. 锁住当前 head.next 地址 [避免 head.next 失联, 假设当前 head 指向 a, head.next 就指向 b]
 *      将 head.next 地址更新为 prev [null -> a -> b 变成 null <- a x b, 此时 a & b 断链]
 *      将 prev 更新为 head 的地址 [此时一轮作业已经完成, 初始的 prev 由 指向 null 变为指向 a]
 *      将 head 更新为开始被锁住的 head.next 地址, 开始下一轮作业 [此时 head 指向 b, prev 指向 a]
 *      即将到来的下一轮就是 锁住 head.next[即c], 更新 head.next 指向 prev[null <- a x b -> c 变为 null <- a <- b x c]
 *      ...
 *   4. 整个链表访问完毕时, prev 作为出参链表的头结点返回[此时 head & next 都已经指向 null]
 */
func reserveLinkedList(head *node) *node {
	var prev, next *node
	for head != nil {
		next = head.next // 锁住当前 head.next 的地址
		head.next = prev // 将指向反转, 此时 head & next 处于失联状态 [但是之前锁住了地址, 故而还能找到该节点]
		prev = head      // 更新 prev 为 head, 代表当前的节点已经反转完毕
		head = next      // 更新 head 为开始锁住的 head.next 重复作业
	}
	return prev
}

/*
 * 双向链表反转[ -> 代表 next 指针, <- 代表 prev 指针]
 *     null      a  ->  b  ->  c -> null <入参: a.prev = nil, c.next = nil>
 *           <-     <-     <-
 *
 *     null  ->  c  ->  b  ->  a -> null <出参: a.next = nil, c.prev = nil>
 *                  <-     <-
 * 解题步骤:
 *   1. 初始化 prev, next 两个节点, 分别代表入参链表当前节点的前一个节点 & 后一个节点
 *   2. 从头部节点开始遍历整个链表, 直至当前节点为 null, 代表整个链表处理完毕
 *   3. 锁住当前 head.next 地址 [避免 head.next 失联, 假设当前 head 指向 a, head.next 就指向 b]
 *      将 head.next 地址更新为 prev [null -> a -> b 变成 null <- a x b, 此时 a & b 断链]
 *      将 head.prev 更新为 next 的地址 [初始的 head.next 由 a 指向 b 编程 a 指向 null, head.prev 由 a 指向 null 指向 b]
 *      将 prev 更新为 head 的地址 [此时一轮作业已经完成, 初始的 prev 由 指向 null 变为指向 a]
 *      将 head 更新为开始被锁住的 head.next 地址, 开始下一轮作业 [此时 head 指向 b, prev 指向 a]
 *      即将到来的下一轮就是 锁住 head.next[即c], 更新 head.next 指向 prev[null <- a x b -> c 变为 null <- a <- b x c]
 *      ...
 *   4. 整个链表访问完毕时, prev 作为出参链表的头结点返回[此时 head & next 都已经指向 null]
		与单链表反转相比, 就多了一步将 head.prev 更新为 head.next[提前锁住] 的步骤
*/
func reserveDoubleLinkedList(head *doubleNode) *doubleNode {
	var prev, next *doubleNode
	for head != nil {
		next = head.next
		head.next = prev
		head.prev = next
		prev = head
		head = next
	}
	return prev
}

/*
 * 删除链表中指定元素值的节点
 * 例如:                       1 -> 2 -> 2 -> 1 -> 3 -> 4 -> 1 -> 5
 *     删除值为 1 的节点, 结果为: 2 -> 2 -> 3 -> 4 -> 5
 *     删除值为 2 的节点, 结果为: 1 -> 1 -> 3 -> 4 -> 1 -> 5
 *     删除值为 3 的节点, 结果为: 1 -> 2 -> 2 -> 1 -> 4 -> 1 -> 5
 *     ...
 * 解题思路:
 *         遍历链表, 直至来到第一个值不等于入参元素值的节点, 从该节点开始处理
 *         锁住当前第一个节点的地址作为返回值, 初始化 prev, cur 两个节点为入参节点
 *         再次遍历链表, 从 cur 节点开始, 直至链表结束
 *         若 cur 节点的值等于目标值, 将 prev.next 更新为 cur.next [丢弃 cur 节点]
 *         若入参节点的值不等于目标值, 将 prev 更新为 cur [记录上一个节点值不是目标值的节点]
 */
func deleteValue(head *node, v int) *node {
	for head != nil {
		if head.v != v {
			break
		}
		head = head.next
	}
	prev, cur := head, head
	for cur != nil {
		if cur.v == v {
			prev.next = cur.next
		} else {
			prev = cur
		}
		cur = cur.next
	}
	return head
}
