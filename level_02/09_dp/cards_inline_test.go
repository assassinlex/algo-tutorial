package _9_dp

import (
	"fmt"
	"testing"
)

func TestWin1(t *testing.T) {
	tests := [][]int{
		{1, 50, 4, 100, 3, 9, 7},
		{20, 5, 14, 100, 33, 19, 67},
		{100, 50, 41, 100, 13, 59, 7},
	}

	for _, test := range tests {
		winner, loser := win1(test)
		if winner < loser {
			winner, loser = loser, winner
		}
		fmt.Printf("赢家: %d  输家 %d\n", winner, loser)
	}
}
