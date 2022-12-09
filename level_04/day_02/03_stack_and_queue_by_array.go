package day_02

import (
	"errors"
	"math"
)

/*
	用数组实现栈 & 队列
	todo:: 实现栈
*/

func NewQueueByArray(limit int) *queueByArray {
	return &queueByArray{
		Arr:   make([]int, limit, limit),
		Begin: 0,
		End:   0,
		Size:  0,
		limit: limit,
	}
}

type queueByArray struct {
	Arr   []int
	Begin int
	End   int
	Size  int
	limit int
}

// Push 入队
func (q *queueByArray) Push(v int) error {
	if q.Size == q.limit {
		return errors.New("队列已满")
	}
	q.Size++
	q.Arr[q.Begin] = v
	q.Begin = q.nextIndex(q.Begin)
	return nil
}

// Pool 出队
func (q *queueByArray) Pool() (int, error) {
	if q.Size == 0 {
		return 0, errors.New("暂无数据")
	}
	q.Size--
	res := q.Arr[q.End]
	q.Arr[q.End] = math.MinInt // 删除弹出的元素, 这一步可有可无, 后续入队的元素会覆盖当前值
	q.End = q.nextIndex(q.End)

	return res, nil
}

// nextIndex 下一个元素该存放的 index
func (q *queueByArray) nextIndex(i int) int {
	res := 0
	if i < q.limit-1 {
		res = i + 1
	}
	return res
}

func (q *queueByArray) IsEmpty() bool {
	return q.Size == 0
}
