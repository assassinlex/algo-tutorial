package day_02

import "errors"

/*
	获取最小栈: 实现一个 GetMin 功能的栈, 即调用该方法时, 返回元素值是当前栈中最小的
		 思路: 用一个结构, 包含两个栈, 数据栈 & 最小值栈
              数据正常的出入栈用数据栈保存, 但是在入栈的过程中, 将数据栈入栈元素的值与最小值栈栈顶的元素值做比较
              若入栈元素值大, 则最小栈将栈顶元素值再次入栈, 否则将入栈元素值入栈
              两个栈的出栈、入栈操作同步[保证两个栈的元素个数相等]
*/

func NewMinStack() *MinStack {
	return &MinStack{
		dataStack: NewStackByList(),
		minStack:  NewStackByList(),
	}
}

type MinStack struct {
	dataStack *stackByList
	minStack  *stackByList
}

// Push 入栈
func (m *MinStack) Push(v int) error {
	if m.minStack.IsEmpty() {
		m.minStack.Push(v)
	} else {
		min, err := m.GetMin()
		if err != nil {
			return err
		}
		if min > v {
			min = v
		}
		m.minStack.Push(min)
	}
	m.dataStack.Push(v)
	return nil
}

func (m *MinStack) Pop() (int, error) {
	if m.dataStack.IsEmpty() {
		return 0, errors.New("暂无数据")
	}
	if _, err := m.minStack.Pop(); err != nil {
		return 0, err
	}
	v, err := m.dataStack.Pop()
	if err != nil {
		return 0, err
	}
	return v, nil
}

func (m *MinStack) GetMin() (int, error) {
	if m.minStack.IsEmpty() {
		return 0, errors.New("暂无数据")
	}
	return m.minStack.Peek()
}

func (m MinStack) IsEmpty() bool {
	return m.dataStack.IsEmpty()
}
