package day_02

import "errors"

/*
	两个栈实现队列
		准备两个栈, 一个入栈[pushStack], 一个出栈[popStack]
		所有的队列 push 操作, 由 pushStack 完成, 所有的队列 poll 操作, 由 popStack 完成
		核心点, 队列的每一次操作, 都需要倒数据: 将 pushStack 的元素全部出栈, 然后将出栈的元素依次入栈到 popStack 中
		注意两点, 每次倒数据必须将 pushStack 中的数据倒完; 倒数据时 popStack 必须不含任何元素[空栈]
	Ps:
		每一次队列的操作都可能导致内部栈所有元素的重新组合, 时间复杂度很高
*/

func NewQueueByStacks() *queueByStacks {
	return &queueByStacks{
		pushStack: NewStackByList(),
		popStack:  NewStackByList(),
	}
}

type queueByStacks struct {
	pushStack *stackByList
	popStack  *stackByList
}

// pushToPop 将 pushStack 中的数据全部倒入 popStack
func (q *queueByStacks) pushToPop() error {
	if q.popStack.IsEmpty() {
		for !q.pushStack.IsEmpty() {
			v, err := q.pushStack.Pop()
			if err != nil {
				return err
			}
			q.popStack.Push(v)
		}
	}
	return nil
}

// Push 入队
func (q *queueByStacks) Push(v int) error {
	q.pushStack.Push(v)
	return q.pushToPop()
}

// Poll 出队
func (q *queueByStacks) Poll() (int, error) {
	if q.pushStack.IsEmpty() && q.popStack.IsEmpty() {
		return 0, errors.New("暂无数据")
	}
	if err := q.pushToPop(); err != nil {
		return 0, nil
	}
	return q.popStack.Pop()
}
