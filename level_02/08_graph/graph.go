package _8_graph

import set "algo/level_02/03_map"

type Graph struct {
	nodes map[int]*Node // 点集
	edges set.MySet     // 边集
}

// NewGraph 构造图
// 参数类型: N * 3 的矩阵, 即所有的边(edge)集[weight, from, to]
// [
//     [5, 0, 7]
//     [3, 0, 1]
// ]
func NewGraph(matrix [][]int) *Graph {
	_graph := &Graph{
		nodes: nil,
		edges: set.NewSet(),
	}

	for _, item := range matrix {
		weight := item[0]
		from := item[1]
		to := item[2]
		if _, ok := _graph.nodes[from]; !ok { // 图添加节点
			_graph.nodes[from] = NewNode(from)
		}
		if _, ok := _graph.nodes[to]; !ok { // 图添加节点
			_graph.nodes[to] = NewNode(to)
		}
		fromNode := _graph.nodes[from]
		toNode := _graph.nodes[to]
		edge := NewEdge(weight, fromNode, toNode)
		fromNode.next = append(fromNode.next, toNode)
		fromNode.out++
		toNode.in++
		fromNode.edge = append(fromNode.edge, edge)
		_graph.edges.Add(edge)
	}

	return _graph
}
