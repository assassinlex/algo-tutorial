package level_02

import (
	"testing"
)

func TestTreeNode_pre(t *testing.T) {
	head := generateTree()
	head.pre()
}

func TestTreeNode_in(t *testing.T) {
	head := generateTree()
	head.in()
}

func TestTreeNode_pos(t *testing.T) {
	head := generateTree()
	head.pos()
}