package _7_binary_tree

import (
	"testing"
)

func TestTreeNode_pre(t *testing.T) {
	head := generateTree()
	head.preRec()
}

func TestTreeNode_in(t *testing.T) {
	head := generateTree()
	head.inRec()
}

func TestTreeNode_pos(t *testing.T) {
	head := generateTree()
	head.posRec()
}