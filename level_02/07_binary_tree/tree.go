package _7_binary_tree

import (
	list "algo/level_02/02_list"
	stack "algo/level_02/05_stack"
	"fmt"
)

// 二叉树

/*
	二叉树的先序、中序、后序遍历

	先序: 任何子树的处理顺序都是. 先头结点、再左子树、后又子树

	中序: 任何子树的处理顺序都是. 先左子树、再头结点、后又子树

	后序: 任何子树的处理顺序都是. 先左子树、再又子树、后头结点

*/

/**
 * 构造二叉树
 *           1
 *       2       3
 *     4   5   6   7
 */
//func generateTree() *_2_list.TreeNode {
//	head := _2_list.NewTreeNode(1)
//	head.Left = _2_list.NewTreeNode(2)
//	head.Right = _2_list.NewTreeNode(3)
//	head.Left.Left = _2_list.NewTreeNode(4)
//	head.Left.Right = _2_list.NewTreeNode(5)
//	head.Right.Left = _2_list.NewTreeNode(6)
//	head.Right.Right = _2_list.NewTreeNode(7)
//	return head
//}

type Tree struct {
	head list.Node
	left *Tree
	right *Tree
}

// 先序 -- 递归
//func (head *_2_list.TreeNode) preRec() {
//	if head == nil {
//		return
//	}
	//fmt.Println(head.Value())
	//head.Left.preRec()
	//head.Right.preRec()
}

//
//// 中序 -- 递归
//func (head *list.TreeNode) inRec() {
//	if head == nil {
//		return
//	}
//	head.Left.inRec()
//	fmt.Println(head.Value())
//	head.Right.inRec()
//}
//
//// 后序 -- 递归
//func (head *list.TreeNode) posRec() {
//	if head == nil {
//		return
//	}
//	head.Left.posRec()
//	head.Right.posRec()
//	fmt.Println(head.Value())
//}

// 先序 -- 压栈
//func (head *list.TreeNode) preByStack() {
//s := stack.NewStack())
//s.Push(head.Value)
//for !s.Empty() {
//	fmt.Println(s.Pop())
//	if head.Right != nil {
//		s.Push(head.Right.Value)
//	}
//	if head.Left != nil {
//		s.Push(head.Left.Value)
//	}
//}
//}
