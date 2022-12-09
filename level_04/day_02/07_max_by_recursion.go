package day_02

import "math"

/*
	获取数组中最大值(递归实现)
	核心: 将数组分为左右两个子数组, 分别处理各自的最大值, 然后将左右两边取大
	     子数组再次分为左右子数组，同样的逻辑进行递归
	     ...
	     直至达到递归的退出条件[base case], 返回当前结果
	Ps: 实际上遍历一遍数组[O(1)]即可解决问题, 但这次是为了研究递归 & master 公式
	    这个例子用 master 公式估算时间复杂度, 结果也是 O(1), 与遍历的时间复杂度一样
*/

func getMax(s []int) int {
	return getMaxProcess(s, 0, len(s)-1)
}

func getMaxProcess(s []int, L, R int) int {
	if L == R { // base case: L & R 是一个数
		return s[L]
	}
	mid := L + ((R - L) >> 1) // 等同于 L + (R - L) / 2
	leftMax := getMaxProcess(s, 0, mid)
	rightMax := getMaxProcess(s, mid+1, R)
	return int(math.Max(float64(leftMax), float64(rightMax)))
}
