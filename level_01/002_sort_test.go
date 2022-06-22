package level_01

import (
	"log"
	"testing"
)

var data = [][]int32 {
	{2, 3, 5, -1, 0, -4, 9, 7},
	{12, 23, 5, -11, 10, -14, 9, 47},
	{-12, 123, 15, -11, 10, -14, 99, 47},
	{10, 9, 8, 7, 6, 5, 4, 3, 2, 1 },
}

// 选择排序
func TestSelectSort(t *testing.T) {
	for _, item := range data {
		log.Println(selectionSort(item))
	}
}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	for _, item := range data {
		log.Println(bubbleSort(item))
	}
}

// 插入排序
func TestInsertSort(t *testing.T) {
	for _, item := range data {
		log.Println(insertSort(item))
	}
}
