package day_02

import "testing"

func TestGetStackMin(t *testing.T) {
	s := []int{1, 1, 2, 3, 3, 2, -1, 0, -2}
	stack := NewMinStack()
	for _, v := range s {
		if err := stack.Push(v); err != nil {
			t.Fatal(err)
		}
	}
	for !stack.IsEmpty() {
		min, err := stack.GetMin()
		_, _ = stack.Pop()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(min) // -2 -1 -1 1 1 1 1 1 1
	}
}
