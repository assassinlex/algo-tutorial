package _9_dp

import (
	"fmt"
	"testing"
)

func TestWay(t *testing.T) {
	var params = []struct {
		M int
		K int
		P int
		N int
	}{
		{1, 3, 4, 6},
		{2, 3, 4, 6},
		{3, 3, 1, 6},
		{4, 3, 2, 6},
		{5, 3, 6, 6},
		{6, 3, 6, 6},
	}

	for _, param := range params {
		fmt.Printf("%v 共有 %d 中方法\n", param, Way1(param.M, param.K, param.P, param.N))
	}

	fmt.Println("===================")

	for _, param := range params {
		fmt.Printf("%v 共有 %d 中方法\n", param, Way2(param.M, param.K, param.P, param.N))
	}
}
