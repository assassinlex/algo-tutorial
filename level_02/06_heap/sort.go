package _6_heap

import array "algo/level_02/01_array"

/**
 * 堆排序
 *
 * 时间复杂度: O(N*logN)
 * 额外空间复杂度: O(1)
 *
 * 将一堆数据[数组]重新排成一个大根堆, 时间复杂度可以优化成为O(N)
 */

// Sort 堆排序, 默认大根堆
// 首先将入参数组调成一个大根堆, 然后将堆顶元素与数组最后一个元素互换位置, 并将数组中是堆的部分大小减一
// 再将新堆做 heapify 操作, 将堆重新排好, 堆顶与数组最后一个元素互换位置, 堆大小减一, 循环作业, 直至对大小变为 0
func Sort(s []int) {
	length := len(s)
	if s == nil || length < 2 {
		return
	}

	h := NewMaxHeap(cap(s))
	for i := 0; i < length; i++ {
		h.heapInsert(s, i) // 大根堆的重排
	}

	length--
	array.Swap(s, 0, length)
	for length > 0 {
		h.heapify(s, 0, length)
		length--
		array.Swap(s, 0, length)
	}
}
