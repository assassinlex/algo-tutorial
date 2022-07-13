package _9_dp

import (
	"fmt"
	"testing"
)

func TestPermutation01(t *testing.T) {
	strings := []string{
		"abc",
		"235",
		"adf",
		"aaa",
	}

	for _, str := range strings {
		fmt.Println(permutation01(str))
	}
}
