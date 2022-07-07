package _8_graph

// Edge 边
type Edge struct {
	weight int   // 权重
	from   *Node //
	to     *Node //
}

// NewEdge 构造边
func NewEdge(weight int, from, to *Node) *Edge {
	return &Edge{
		weight: weight,
		from:   from,
		to:     to,
	}
}
