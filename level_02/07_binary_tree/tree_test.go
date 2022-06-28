package _7_binary_tree

import "testing"

func TestTreeNode_pre(t *testing.T) {
	tree := generateTree()
	preRec(tree)
}

func TestTreeNode_in(t *testing.T) {
	head := generateTree()
	inRec(head)
}

func TestTreeNode_pos(t *testing.T) {
	head := generateTree()
	posRec(head)
}

func TestTreeNode_preStack(t *testing.T) {
	head := generateTree()
	preByStack(head)
}

func TestTreeNode_inStack(t *testing.T) {
	head := generateTree()
	inByStack(head)
}

func TestTreeNode_posStack(t *testing.T) {
	head := generateTree()
	posByStack(head)
}
