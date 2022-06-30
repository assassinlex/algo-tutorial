package _3_map

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	set := NewSet(10)

	for i := 0; i < 5; i++ {
		set.Add(i)
	}

	for v := range set {
		t.Log(v)
	}
}
