package _8_graph

// 默认集合大小
var size = 10

// Node 顶点[节点]
type Node struct {
	id   int     // 编号
	in   int     // 入度
	out  int     // 出度
	next []*Node //
	edge []*Edge //
}

// NewNode 构造顶点
func NewNode(id int) *Node {
	return &Node{
		id:   id,
		in:   0,
		out:  0,
		next: make([]*Node, 0, size),
		edge: make([]*Edge, 0, size),
	}
}
