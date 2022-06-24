package _4_queue

import (
	list "algo/level_02/02_list"
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {

	q := NewQueue()

	for i := 0; i < 5; i++ {
		q.Enqueue(list.NewListNode(int64(i)))
	}

	for q.Len() > 0 {
		fmt.Printf("出队前长度: %d, 出队元素: %d, 队列后长度: %d\n", q.Len(), q.Dequeue().Value(), q.Len())
	}
}
