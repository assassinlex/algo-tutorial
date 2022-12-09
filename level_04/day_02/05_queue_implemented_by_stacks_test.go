package day_02

import "testing"

func TestQueueByStacks(t *testing.T) {
	queue := NewQueueByStacks()
	s := []int{1, 2, 3, 4, 5}
	for _, v := range s {
		if err := queue.Push(v); err != nil {
			t.Fatal(err)
		}
	}
	for i := 0; i < len(s); i++ {
		v, err := queue.Poll()
		if err != nil {
			t.Fatal(err)
		}
		t.Log(v)
	}
}
