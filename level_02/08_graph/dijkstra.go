package _8_graph

import (
	set "algo/level_02/03_map"
	"math"
)

/**
 * Dijkstra [/ˈdikstrɑ/ 或 /ˈdɛikstrɑ/, 最短路算法之一]
 *
 * 定义[维基百科]:
 *      由荷兰计算机科学家 艾兹赫尔·戴克斯特拉 在1956年发现的算法.
 *      Dijkstra算法 使用类似广度优先搜索的方法解决赋权图的单源最短路径问题.
 *      Dijkstra算法 原始版本仅适用于找到两个顶点之间的最短路径, 后来更常见的变体固定了一个顶点作为源结点,
 *      			然后找到该顶点到图中所有其它结点的最短路径，产生一个最短路径树.
 *
 * 注意:
 *      如果图中的顶点表示城市, 而边上的权重表示城市间开车行经的距离, 该算法可以用来找到两个城市之间的最短路径.
 *      绝大多数的 Dijkstra算法 不能有效处理带有负权边的图.
 */

// dijkstra01 提供一个随机固定顶点, 返回这个顶点到所有其他顶点的所需最小权重表
func dijkstra01(start *Node) map[*Node]int {
	weights := make(map[*Node]int)                              // 从 start 顶点到图中其他所有顶点的权重表
	weights[start] = 0                                          // 标记 start 到 start 边权重为 0
	lockedNodes := set.NewSet(size)                             // 标记已经处理过的顶点
	node := getMinWeightAndUnselectedNode(weights, lockedNodes) // start 的最小权重邻点
	if node != nil {
		weight := weights[node]          // 当前最小权重邻点的权重值
		for _, edge := range node.edge { // 迭代当前最小权重邻点的所有边[边的另一端是其他顶点]
			if _, ok := weights[edge.to]; ok { // 更新到当前边的下一个跳转点最小权重记录
				weights[edge.to] = int(math.Min(float64(weights[edge.to]), float64(weight+edge.weight)))
			} else {
				weights[edge.to] = weight + edge.weight
			}
		}
		lockedNodes.Add(node)                                      // 锁定当前最小权重顶点[即不在更新到达当前点的最小权重值]
		node = getMinWeightAndUnselectedNode(weights, lockedNodes) // 更新所有边的邻点中的最小权重顶点
	}
	return weights
}

// getMinWeightAndUnselectedNode 获取当前顶点的最小权重且未被锁定的邻点
func getMinWeightAndUnselectedNode(weights map[*Node]int, lockNodes set.MySet) *Node {
	var minWeightNode *Node
	var minWeight = math.MaxInt
	for node, weight := range weights {
		if !lockNodes.Contain(node) && weight < minWeight { // 跳转点未被锁定且跳转所需权重小于之前到达跳转点的权重记录时更新
			minWeightNode = node
			minWeight = weight
		}
	}
	return minWeightNode
}
