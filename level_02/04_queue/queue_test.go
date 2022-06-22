package _4_queue

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {

	q := NewQueue([]int64 {
		1, 2, 3, 4, 5, 6,
	})

	for q.Len() > 0 {
		fmt.Printf("获取当前头元素: %d, 队列长度: %d\n", q.Pop(), q.Len())
	}
}
