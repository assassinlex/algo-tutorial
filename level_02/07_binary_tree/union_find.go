package _7_binary_tree

import (
	list "algo/level_02/02_list"
	stack "algo/level_02/05_stack"
)

/**
 * 并查集 [Disjoint-set data structure]
 *
 * 定义[维基百科]:
 *      一种数据结构[树形结构], 用于处理一些不交集(Disjoint sets，一系列没有重复元素的集合)的合并及查询问题
 *      因为支持查询和合并两种操作, 并查集也常被称为联合-查找数据结构(Union-find data structure),
 *      或者联合-查找集合(Merge-find set)
 *
 * 支持方法:
 *      查询: 查询某个袁术属于哪个集合, 通常返回集合内的代表元素[一般用于判别两个元素是否属于同一个集合]
 *      合并: 将两个集合合并
 *      添加: 添加一个新的集合[只包含新元素本身]. 此方法基本被忽略, 故常见方法一般代指以上两种
 */

type DisjointSet interface {
	Union(v1, v2 int64)
	Find(v1 int64) int64
}

type MyDisjointSet struct {
	nodes   map[int64]list.Node     // 节点表: 元素值对应的节点
	fathers map[list.Node]list.Node // 父节点表: 节点对应的父节点
	size    map[list.Node]int64     // 代表元素所在集合对应的元素个数大小
}

// Find 查找代表元素
func (m *MyDisjointSet) Find(v1 int64) int64 {
	return m.findFather(m.nodes[v1]).Value()
}

// Union 合并两个集合[参数是两个集合各自的代表元素]
func (m *MyDisjointSet) Union(v1, v2 int64) {
	v1Head := m.findFather(m.nodes[v1])
	v2Head := m.findFather(m.nodes[v2])
	if v1Head != v2Head { // 当两个元素的代表元素节点不是同一个节点时[即在两个不同的集合]
		if m.size[v1Head] <= m.size[v2Head] { // 合并规则: 小集合挂到大集合上, 更新大集合代表元素大小, 移除小代表元素大小
			m.fathers[v1Head] = v2Head
			m.size[v2Head] += m.size[v1Head]
			delete(m.size, v1Head)
		} else {
			m.fathers[v2Head] = v1Head
			m.size[v1Head] += m.size[v2Head]
			delete(m.size, v2Head)
		}
	}
}

// 寻找代表元素节点
func (m *MyDisjointSet) findFather(cur list.Node) list.Node {
	path := stack.NewStack()
	for cur != m.fathers[cur] { // 当前节点元素与自己的父元素不是用一个对象 [即不是代表元素(父元素指向自己)]
		path.Push(cur)       // 入栈[保证最后栈顶是代表元素]
		cur = m.fathers[cur] // 更新当前元素为其父元素继续作业
	}
	for path.Size() > 0 {
		m.fathers[path.Pop()] = cur // 将 cur 所在的这个集合中的所有元素的父节点更新为 cur[代表元素节点]
	}
	return cur
}

// 两个元素是否属于同一个集合
func (m *MyDisjointSet) isSameSet(v1, v2 int64) bool {
	return m.findFather(m.nodes[v1]) == m.findFather(m.nodes[v2])
}

// Size 获取并查集代表元素个数
func (m *MyDisjointSet) Size() int {
	return len(m.size)
}
