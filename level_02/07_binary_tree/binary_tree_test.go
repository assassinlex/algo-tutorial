package _7_binary_tree

import (
	list "algo/level_02/02_list"
	"testing"
)

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

func TestBFS_bfs(t *testing.T) {
	head := generateTree()
	bfs(head)
}

func TestBFS_GetTreeMaxWidthWithMap(t *testing.T) {
	head := generateTree()
	t.Log(GetTreeMaxWidthWithMap(head))
}

func TestBFS_GetTreeMaxWidthWithoutMap(t *testing.T) {
	head := generateTree()
	t.Log(GetTreeMaxWidthWithoutMap(head))
}

/**
 * 两种类型的二叉树
 *         1          1
 *      2                 3
 *         5          6
 */
func TestSerialize_PreSerialize(t *testing.T) {
	tree := list.NewTreeNode(1)
	tree.SetLeft(list.NewTreeNode(2))
	tree.Left().SetRight(list.NewTreeNode(5))
	s := PreSerialize(tree)
	t.Log(s)
	for i := 0; i < len(s); i++ {
		t.Log(s[i])
	}

	tree = list.NewTreeNode(1)
	tree.SetRight(list.NewTreeNode(3))
	tree.Right().SetLeft(list.NewTreeNode(6))
	s = PreSerialize(tree)
	t.Log(s)
	for i := 0; i < len(s); i++ {
		t.Log(s[i])
	}
}
