package _9_dp

import "fmt"

/**
 * 汉诺塔:
 *   有三根杆子A, B, C, A杆上有 N 个 (N>1) 穿孔圆盘，盘的尺寸由下到上依次变小. 要求按下列规则将所有圆盘移至 C 杆:
 *      1. 每次只能移动一个圆盘
 *      2. 大盘不能叠在小盘上面
 *
 *    可以将 3 根杆子理解为 left, mid, right, 或者是 from, to, other
 */

// hanoi01 原始版
func hanoi01(n int) {
	leftToRight(n)
}

// leftToMid 把 1~N 层圆盘从 left 移动到 mid
// 先把 1~N-1 从 left  移动到 right [继续递归]
// 再把 N     从 left  移动到 mid   [函数执行: 输出内容]
// 最后 1~N-1 从 right 移动到 mid   [继续递归]
func leftToMid(n int) {
	if n == 1 {
		fmt.Println("move 1 from left to mid")
		return
	}
	leftToRight(n - 1)
	fmt.Printf("move %d from left to mid\n", n)
	rightToMid(n - 1)
}

// leftToRight 把 1~N 层圆盘从 left 移动到 right
// 先把 1~N-1 从 left  移动到 mid   [继续递归]
// 再把 N     从 left  移动到 right [函数执行: 输出内容]
// 最后 1~N-1 从 mid   移动到 right [继续递归]
func leftToRight(n int) {
	if n == 1 {
		fmt.Println("move 1 from left to right")
		return
	}
	leftToMid(n - 1)
	fmt.Printf("move %d from left to right\n", n)
	midToRight(n - 1)
}

// midToLeft 把 1~N 层圆盘从 mid 移动到 left
// 先把 1~N-1 从 mid   移动到 right [继续递归]
// 再把 N     从 mid   移动到 left  [函数执行: 输出内容]
// 最后 1~N-1 从 right 移动到 left  [继续递归]
func midToLeft(n int) {
	if n == 1 {
		fmt.Println("move 1 from mid to left")
		return
	}
	midToRight(n - 1)
	fmt.Printf("move %d from mid to left\n", n)
	rightToLeft(n - 1)
}

// midToRight 把 1~N 层圆盘从 mid 移动到 right
// 先把 1~N-1 从 mid   移动到 left  [继续递归]
// 再把 N     从 mid   移动到 right [函数执行: 输出内容]
// 最后 1~N-1 从 left  移动到 right [继续递归]
func midToRight(n int) {
	if n == 1 {
		fmt.Println("move 1 from mid to right")
		return
	}
	midToLeft(n - 1)
	fmt.Printf("move %d from mid to right\n", n)
	leftToRight(n - 1)
}

// rightToLeft 把 1~N 层圆盘从 right 移动到 left
// 先把 1~N-1 从 right 移动到 mid   [继续递归]
// 再把 N     从 right 移动到 left  [函数执行: 输出内容]
// 最后 1~N-1 从 mid   移动到 left  [继续递归]
func rightToLeft(n int) {
	if n == 1 {
		fmt.Println("move 1 from right to left")
		return
	}
	rightToMid(n - 1)
	fmt.Printf("move %d from right to left\n", n)
	midToLeft(n - 1)
}

// rightToMid 把 1~N 层圆盘从 right 移动到 mid
// 先把 1~N-1 从 right 移动到 left  [继续递归]
// 再把 N     从 right 移动到 mid   [函数执行: 输出内容]
// 最后 1~N-1 从 left  移动到 mid   [继续递归]
func rightToMid(n int) {
	if n == 1 {
		fmt.Println("move 1 from right to mid")
		return
	}
	rightToLeft(n - 1)
	fmt.Printf("move %d from right to mid\n", n)
	leftToMid(n - 1)
}

// hanoi2 正常版
func hanoi2(n int) {
	if n > 0 {
		move(n, "left", "right", "mid")
	}
}

// move 将 1~N 层圆盘从 from 移动到 to, other 为临时存放区
// 先把 1~N-1 从 from  移动到 other [继续递归]
// 再把 N     从 from  移动到 to    [函数执行: 输出内容]
// 最后 1~N-1 从 other 移动到 to    [继续递归]
func move(n int, from, to, other string) {
	if n == 1 {
		fmt.Printf("move 1 from %s to %s\n", from, to)
		return
	}
	move(n-1, from, other, to) // 从 from 移动到 other[参数: from = from, to = other, other = to]
	fmt.Printf("move %d from %s to %s\n", n, from, to)
	move(n-1, other, to, from) // 从 other 移动到 to[参数: from = other, to = to, other = from]
}
