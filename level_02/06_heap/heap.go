package _6_heap

import (
	"errors"
	"fmt"
)

/**
 * 堆结构[heap]
 *
 * 定义:
 *     堆是一种由数组实现的完全二叉树结构[逻辑上是一颗完全二叉树, 实际由数组存储]
 *     根据每棵子树的堆顶的值的大小, 值是子树中的最大值即为大根堆, 最小值为小根堆
 *
 * 节点的索引公式:
 *     父亲节点: (i - 1) / 2 [变种: i / 2        eq: i >> 2]
 *     子左节点: (i * 2) + 1 [变种: i * 2        eq: i << 2]
 *     子右节点: (i * 2) + 2 [变种: i * 2 + 1    eq: (i << 2) | 1]
 *
 * Ps: 有一些堆的实现的索引是基于变种来的, 因为变种的位运算效率高于一般的四则运算
 */

// MaxHeap 大根堆
type MaxHeap struct {
	heap     []int
	size     int
	capacity int
}

// NewMaxHeap 构造大根堆
func NewMaxHeap(capacity int) *MaxHeap {
	return &MaxHeap{
		heap:     make([]int, 0, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (h *MaxHeap) IsEmpty() bool {
	return h.size == 0
}

func (h *MaxHeap) IsFull() bool {
	return h.size == h.capacity
}

func (h MaxHeap) Println() {
	fmt.Print("[ ")
	for i := 0; i < h.size; i++ {
		fmt.Printf("%d ", h.heap[i])
	}
	fmt.Print("]\n")
}

// Push 添加元素
func (h *MaxHeap) Push(v int) error {
	if h.IsFull() {
		return errors.New("堆已满")
	}
	//h.heap[h.size] = v // 添加元素
	h.heap = append(h.heap, v) // 添加元素
	h.heapInsert()             // 更新数组结构
	h.size++                   // 更新堆大小
	return nil
}

// Pop 弹出堆顶元素[最大值]
func (h *MaxHeap) Pop() (int, error) {
	if h.IsEmpty() {
		return 0, errors.New("堆为空")
	}
	ret := h.heap[0]                                      // 锁住要返回的值
	h.size--                                              // 更新堆大小
	h.heap[0], h.heap[h.size] = h.heap[h.size], h.heap[0] // 数组首尾互换
	h.heapify()
	return ret, nil
}

// heapInsert
// 新添加的元素, 默认在数组的最后一个位置, 然后进行大根堆重排
// 新元素作为二叉树节点与自己的父亲节点比大小,
// 若新元素大则父亲节点交换位置[更新索引], 循环作业; 否则就保持原位置[索引不变]
func (h *MaxHeap) heapInsert() {
	curIdx := h.size
	for h.heap[curIdx] > h.heap[(curIdx-1)/2] {
		h.heap[curIdx], h.heap[(curIdx-1)/2] = h.heap[(curIdx-1)/2], h.heap[curIdx]
		curIdx = (curIdx - 1) / 2
	}
}

// heapify
// 从当前根父亲节点开始, 比较父亲节点与子节点的大小
// 若父亲节点比子节点小, 将父亲节点下沉,
// 直至左右节点较大的子节点不再比父亲节点大, 或者父亲节点不在有子节点
// Ps: 此方法有个前提是将最后一个子节点放到了父亲节点的位置, 保证首次作业时父亲节点一定比子节点小
func (h *MaxHeap) heapify() {
	parentIdx := 0
	left := parentIdx*2 + 1 // 当前节点左子结点索引
	for left < h.size {     // 当前节点有子左节点 [h.size 代表堆中最后一个元素的索引位置]
		largest := left // 默认当前较大的节点索引为左子节点索引
		if left+1 < h.size && h.heap[left] < h.heap[left+1] {
			largest = left + 1 // 若有子右节点[left+1], 且子右节点值较大
		}
		if h.heap[largest] <= h.heap[parentIdx] {
			largest = parentIdx // 若当前子节点中较大值已小于父亲节点, 将父亲节点更新, 结束循环
		}
		if largest == parentIdx { // 若当前最大子节点就是父亲节点本身, 代表已无后续子节点
			break
		}
		h.heap[largest], h.heap[parentIdx] = h.heap[parentIdx], h.heap[largest]
		parentIdx = largest
		left = parentIdx*2 + 1
	}
}
