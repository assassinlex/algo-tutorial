package _7_binary_tree

import (
	list "algo/level_02/02_list"
	queue "algo/level_02/04_queue"
	"fmt"
	"math"
)

/**
 * Breadth-First Search [BFS]
 * 广度优先搜索算法、广度优先遍历、宽度优先遍历、层次遍历
 * 从二叉树头结点开始, 从上至下依次按层遍历, 在每一层中从左往右遍历
 *
 * 常见题型: 获取二叉树的最大宽度
 */

// 广度优先算法简单案例
func bfs(head list.Node) {
	if head == nil {
		return
	}

	// 先让头结点入队, 保证队列非空, 然后就可以开始循环作业
	q := queue.NewQueue()
	q.Enqueue(head)

	for q.Len() > 0 {
		// 出队就输出当前头结点
		head = q.Dequeue()
		fmt.Println(head.Value())
		// 有左子树就让左子树入队
		if head.Left() != nil {
			q.Enqueue(head.Left())
		}
		// 有右子树就让右子树入队
		if head.Right() != nil {
			q.Enqueue(head.Right())
		}
	}
}

// GetTreeMaxWidthWithMap 获取二叉树最大宽度 -- 用 map 版本
func GetTreeMaxWidthWithMap(head list.Node) int {
	max := 0
	if head == nil {
		return max
	}

	// 存放二叉树节点
	q := queue.NewQueue()
	// 存放节点所在层级
	m := make(map[list.Node]int)
	// 当前统计层级 [约定从树顶点开始统计, 故初始化为 1]
	curLevel := 1
	// 当前层级节点数量 [约定以节点出队时更新, 故初始化为 0]
	curLevelNodeCnt := 0
	// 节点入队, 然后开始循环作业
	q.Enqueue(head)
	// 记录节点所在层级
	m[head] = curLevel

	for q.Len() > 0 {
		head = q.Dequeue()
		// 临时记录当前节点所在层级
		curNodeLevel := m[head]
		// 有左子树将左子树入队, 并记录节点所在层级 [父节点所在层级 + 1]
		if head.Left() != nil {
			m[head.Left()] = curNodeLevel + 1
			q.Enqueue(head.Left())
		}
		// 有右子树将右子树入队, 并记录节点所在层级
		if head.Right() != nil {
			m[head.Right()] = curNodeLevel + 1
			q.Enqueue(head.Right())
		}
		// 若当前节点所在层级与当前统计层级相等, 代表当前层级可能还未统计完成,
		// 当前节点可能是, 也可能不是当前层级最后一个节点
		if curNodeLevel == curLevel {
			// 更新当前层级节点数
			curLevelNodeCnt++
		} else {
			// 一旦第一次发生当前节点所在层级与不相等时, 当前层级一定遍历完全,
			// 因为此时的元素必定为下一层级的第一个节点,
			// 开始了下一层级的统计, 此时更新全局最大节点数
			max = int(math.Max(float64(max), float64(curLevelNodeCnt)))
			// 开始统计下一层级
			curLevel++
			// 此时因队列已出队下一层级的一个元素, 故下一层级节点数初始化为 1 而不是 0
			curLevelNodeCnt = 1
		}
	}

	// 因 max 的更新时机是当前层级统计完成, 新层(下一层来临之际)
	// 故最后一层节点数的值永远不会与全局 max 比较来更新 max
	// 所以树遍历完成后需要在做一轮 max 的更新
	max = int(math.Max(float64(max), float64(curLevelNodeCnt)))

	return max
}

// GetTreeMaxWidthWithoutMap 获取二叉树最大宽度 -- 不用 map 版本
func GetTreeMaxWidthWithoutMap(head list.Node) int {
	max := 0

	if head == nil {
		return max
	}

	q := queue.NewQueue()
	q.Enqueue(head)

	// 当前统计层级的最右节点(最后一个节点, 初始化为树的顶点)
	curLevelEndNode := head
	// 下一层级的最右节点(因不知道是哪个节点, 初始化为 nil)
	var nextLevelEndNode list.Node
	// 当前统计层级节点数
	curLevelNodeCnt := 0

	for q.Len() > 0 {
		curNode := q.Dequeue()
		// 处理当前队列元素的左右子树, 并更新下一层级的最右节点
		if curNode.Left() != nil {
			q.Enqueue(curNode.Left())
			nextLevelEndNode = curNode.Left()
		}
		if curNode.Right() != nil {
			q.Enqueue(curNode.Right())
			nextLevelEndNode = curNode.Right()
		}
		// 更新当前统计层级节点数
		curLevelNodeCnt++
		// 当 当前节点与当前层级最右节点是同一个对象时代表当前层级应该结束
		if curNode == curLevelEndNode {
			max = int(math.Max(float64(max), float64(curLevelNodeCnt)))
			// 重置当前统计层级的节点个数
			curLevelNodeCnt = 0
			// 开启下一层级的统计(更新层级的最右节点为此时下一层的最右节点)
			curLevelEndNode = nextLevelEndNode
		}
	}

	return max
}
