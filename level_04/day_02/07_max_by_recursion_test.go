package day_02

import "testing"

func TestGetMax(t *testing.T) {
	collection := [][]int{
		{1, 2, 3, 4, 5},
		{1, 2, 3, 4, -5},
		{1, 2, 3, -4, -5},
		{1, 2, -3, -4, -5},
		{1, -2, -3, -4, -5},
		{-1, 2, -3, -4, -5},
	}
	for _, item := range collection {
		t.Log(getMax(item))
	}
}
