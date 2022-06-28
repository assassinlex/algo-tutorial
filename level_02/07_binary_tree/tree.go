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

type Tree = list.TreeNode

func generateTree() *Tree {
	tree := list.NewTreeNode(1)
	tree.SetLeft(list.NewTreeNode(2))
	tree.SetRight(list.NewTreeNode(3))
	tree.Left().SetLeft(list.NewTreeNode(4))
	tree.Left().SetRight(list.NewTreeNode(5))
	tree.Right().SetLeft(list.NewTreeNode(6))
	tree.Right().SetRight(list.NewTreeNode(7))
	return tree
}

// 递归序
// 每个元素都有 3 次到达的机会
// 1. 第一次到达元素本身  2. 从左子树返回  3. 从右子树返回
// 若是每一次都输出结果, 整个打印顺序是:
// 1 2 4 4 4 2 5 5 5 2 1 3 6 6 6 3 7 7 7 3 1
// 先序: 在元素第一次出现就输出
// 中序: 在元素第二次出现就输出
// 后序: 在元素第三次出现就输出

// 先序 -- 递归
func preRec(head list.Node) {
	if head == nil {
		return
	}
	fmt.Println(head.Value())
	preRec(head.Left())
	preRec(head.Right())
}

// 中序 -- 递归
func inRec(head list.Node) {
	if head == nil {
		return
	}
	inRec(head.Left())
	fmt.Println(head.Value())
	inRec(head.Right())
}

// 后序 -- 递归
func posRec(head list.Node) {
	if head == nil {
		return
	}
	posRec(head.Left())
	posRec(head.Right())
	fmt.Println(head.Value())
}

// 先序 -- 压栈
// 首先入栈头结点, 保证栈内有一个元素, 然后循环作业
// 出栈就输出元素, 有右子树就入栈右子树, 有左子树就入栈左子树
// 出栈顺序就是: 头结点  ->  左子树  ->  右子树 [先序遍历]
func preByStack(head list.Node) {
	s := stack.NewStack()
	s.Push(head)
	for s.Size() > 0 {
		head = s.Pop()
		fmt.Println(head.Value())
		if head.Right() != nil {
			s.Push(head.Right())
		}
		if head.Left() != nil {
			s.Push(head.Left())
		}
	}
}

// 中序 -- 压栈
// 解题规则:
// 1. 左子树(包括头结点)依次入栈, 直到左子树完全入栈
// 2. 出栈就输出元素, 并将弹出元素的右子树重复步骤 1 [将右子树的左子树依次入栈]
// 结束作业的条件是 栈为空 || 节点完全遍历
func inByStack(head list.Node) {
	s := stack.NewStack()
	for s.Size() > 0 || head != nil {
		if head != nil {
			s.Push(head)
			head = head.Left()
		} else {
			head = s.Pop()
			fmt.Println(head.Value())
			head = head.Right()
		}
	}
}

// 后序 -- 压栈
// 首先入栈头结点, 保证栈内有一个元素, 然后循环作业
// 出栈就输出元素, 有左子树就入栈左子树, 有右子树就入栈右子树
// 出栈顺序就是: 头结点  ->  右子树  ->  左子树
//
// 重点:
// 此时做出一点改动, 额外生成另一个栈, 将出栈就输出元素改为入栈[新生成的栈]
// 那么新栈的出栈顺序: 左子树  ->  右子树  ->  头结点 [后序遍历]
func posByStack(head list.Node) {
	s := stack.NewStack()
	s2 := stack.NewStack()
	s.Push(head)
	for s.Size() > 0 {
		head = s.Pop()
		s2.Push(head)
		if head.Left() != nil {
			s.Push(head.Left())
		}
		if head.Left() != nil {
			s.Push(head.Right())
		}
	}

	for s2.Size() > 0 {
		fmt.Println(s2.Pop().Value())
	}
}
