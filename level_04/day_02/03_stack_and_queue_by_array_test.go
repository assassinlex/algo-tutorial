package day_02

import "testing"

func TestArrayQueue(t *testing.T) {
	q := NewQueueByArray(5)
	s := [...]int{1, 2, 3, 4, 5}
	for _, v := range s {
		if err := q.Push(v); err != nil {
			t.Fatal(err)
		}
	}
	for !q.IsEmpty() {
		v, _ := q.Pool()
		t.Log(v)
	}
}
