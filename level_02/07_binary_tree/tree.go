package _7_binary_tree

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
	Value int64
	Left  *TreeNode
	Right *TreeNode
}

/**
 * 构造二叉树
 *           1
 *       2       3
 *     4   5   6   7
 */
func generateTree() *TreeNode {
	head := &TreeNode{Value: 1}
	head.Left = &TreeNode{Value: 2}
	head.Right = &TreeNode{Value: 3}
	head.Left.Left = &TreeNode{Value: 4}
	head.Left.Right = &TreeNode{Value: 5}
	head.Right.Left = &TreeNode{Value: 6}
	head.Right.Right = &TreeNode{Value: 7}
	return head
}

// 先序 -- 递归
func (head *TreeNode) preRec() {
	if head == nil {
		return
	}
	fmt.Println(head.Value)
	head.Left.preRec()
	head.Right.preRec()
}

// 中序 -- 递归
func (head *TreeNode) inRec() {
	if head == nil {
		return
	}
	head.Left.inRec()
	fmt.Println(head.Value)
	head.Right.inRec()
}

// 后序 -- 递归
func (head *TreeNode) posRec() {
	if head == nil {
		return
	}
	head.Left.posRec()
	head.Right.posRec()
	fmt.Println(head.Value)
}

// 先序 -- 压栈
func (head *TreeNode) preByStack() {
	//s := stack.NewStack([]*list.Node{})
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
}
