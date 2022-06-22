// 1. 一些常见的排序算法
// 2. 前缀和
// 2. 对数器

package level_01

import "math/rand"

// 选择排序
func selectionSort(s []int32) []int32 {
	// 找出最小值所在索引
	// 0 ~ n-1
	// 1 ~ n-1
	// 2 ~ n-1
	n := len(s)
	if s == nil || n < 2 {
		return s
	}
	for i := 0; i < n; i++ {
		minValueIdx := i
		// 从 i 开始，逐步往后比较最小值，若发现更小值，更新最小值索引，交换索引对应的值
		// 内层循环的主要目的是找到当前未排序的序列中最小值所在索引
		for j := i+1; j < n; j++ {
			if s[j] < s[minValueIdx] {
				minValueIdx = j
			}
		}
		// 外层循环在找到未排序的序列中最小值后，固定索引，将最小值排好序
		s[i], s[minValueIdx] = s[minValueIdx], s[i]
	}

	return s
}


// bubbleSort 冒泡排序
func bubbleSort(s []int32) []int32 {
	// 0 ~ n-1
	// 0 ~ n-2
	// 0 ~ n-3
	n := len(s)
	if s == nil || n < 2 {
		return s
	}

	for i := n - 1; i >= 0 ; i-- {
		// 内层循环目的是将未排序的序列, 相邻的两个数两两比较，大值向右移动，最终一个循环将序列中最大值移到最右边
		for j := 1; j <= i; j++ {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}

	return s
}

// insertSort 插入排序
func insertSort(s []int32) []int32 {
	// 1 ~ 2
	// 2 ~ 3
	// 3 ~ n-1
	n := len(s)
	if s == nil || n < 2 {
		return s
	}

	// 将外层循环的每个索引看做插入点
	for i := 1; i < n; i++ {
		// i ~ i-1
		// 2 ~ 1
		// 1 ~ 0
		// 当插入点的值小于插入点左边第一个的值，交换值，插入点往左边移动一位
		for j := i; j > 0; j-- {
			if s[j] < s[j-1] {
				s[j], s[j-1] = s[j-1], s[j]
			}
		}
	}

	return s
}

// 返回 [0, 10) 之间的数的概率(等概率)
func equalP(times, n int) float32 {
	count := 0
	for i := 0; i < times; i++ {
		if rand.Intn(10) < n {
			count++
		}
	}

	return float32(count) / float32(times)
}
