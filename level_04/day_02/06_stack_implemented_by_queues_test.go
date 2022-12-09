package day_02

import (
	"math/rand"
	"testing"
)

// 对数器
func TestStackByQueues(t *testing.T) {
	stack1 := NewStackByQueues()
	stack2 := NewStackByList()
	times, max := 1000000, 1000000
	for i := 0; i < times; i++ {
		v := rand.Float32()
		if stack1.IsEmpty() {
			if !stack2.IsEmpty() {
				t.Fatal("Oops")
			}
			_v := int(rand.Float32() * float32(max))
			stack1.Push(_v)
			stack2.Push(_v)
		} else {
			if v < 0.25 {
				_v := int(rand.Float32() * float32(max))
				stack1.Push(_v)
				stack2.Push(_v)
			} else if v < 0.5 {
				v1, err1 := stack1.Peek()
				v2, err2 := stack2.Peek()
				if err1 != nil || err2 != nil {
					t.Fatal(err1, err2)
				}
				if v1 != v2 {
					t.Fatal("Oops")
				}
			} else if v < 0.75 {
				v1, err1 := stack1.Pop()
				v2, err2 := stack2.Pop()
				if err1 != nil || err2 != nil {
					t.Fatal(err1, err2)
				}
				if v1 != v2 {
					t.Fatal("Oops")
				}
			} else {
				if stack1.IsEmpty() != stack2.IsEmpty() {
					t.Fatal("Oops")
				}
			}
		}
	}
}
