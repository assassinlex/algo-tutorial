package level_02

import (
	"fmt"
)

// 二叉树

/*
	二叉树的先序、中序、后序遍历

	先序: 任何子树的处理顺序都是. 先头结点、再左子树、后又子树

	中序: 任何子树的处理顺序都是. 先左子树、再头结点、后又子树

	后序: 任何子树的处理顺序都是. 先左子树、再又子树、后头结点

 */

type TreeNode struct {
	Value int
	Left *TreeNode
	Right *TreeNode
}

// 构造二叉树
func generateTree() *TreeNode {
	head := &TreeNode{Value: 1}
	head.Left = &TreeNode{Value: 3}
	head.Right = &TreeNode{Value: 4}
	head.Left.Left = &TreeNode{Value: 2}
	head.Left.Right = &TreeNode{Value: 5}
	head.Right.Left = &TreeNode{Value: 3}
	head.Right.Right = &TreeNode{Value: 7}
	return head
}

// 先序
func (head *TreeNode) pre() {
	if head == nil {
		return
	}
	fmt.Println(head.Value)
	head.Left.pre()
	head.Right.pre()
}

// 中序
func (head *TreeNode) in() {
	if head == nil {
		return
	}
	head.Left.in()
	fmt.Println(head.Value)
	head.Right.in()
}

// 后序
func (head *TreeNode) pos() {
	if head == nil {
		return
	}
	head.Left.pos()
	head.Right.pos()
	fmt.Println(head.Value)
}

// 宽度优先遍历 [tips: 用队列结构]
