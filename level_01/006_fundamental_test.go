package level_01

import (
	"log"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := [][]int{
		{1, 2, 3},
		{11, 12, 23},
		{10, 20, 30},
	}

	for _, test := range tests {
		res := add(test[0], test[1])
		if res != test[2] {
			log.Fatalf("except %d, get %d", test[2], res)
		}
	}
}

func TestMinus(t *testing.T) {
	tests := [][]int{
		{1, 2, -1},
		{11, 12, -1},
		{10, 20, -10},
		{30, 20, 10},
	}

	for _, test := range tests {
		res := minus(test[0], test[1])
		if res != test[2] {
			log.Fatalf("except %d, get %d", test[2], res)
		}
	}
}

func TestMulti(t *testing.T) {
	tests := [][]int{
		{1, 2, 2},
		{11, 2, 22},
		{10, 20, 200},
		{30, 20, 600},
	}

	for _, test := range tests {
		res := multi(test[0], test[1])
		if res != test[2] {
			log.Fatalf("except %d, get %d", test[2], res)
		}
	}
}
