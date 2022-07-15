package _9_dp

import (
	"fmt"
	"testing"
)

func TestQueen1(t *testing.T) {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	}

	for _, num := range nums {
		fmt.Printf("%d 皇后的可行解: %d\n", num, Queen1(num))
	}
}
func TestQueen2(t *testing.T) {
	nums := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14,
	}

	for _, num := range nums {
		fmt.Printf("%d 皇后的可行解: %d\n", num, Queen2(num))
	}
}
