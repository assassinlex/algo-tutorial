package day_01

import "testing"

var ss = [][]int{
	{1},
	{3, 5, 4, 6, 1},
	{3, 5, 2, 6, 12, 55},
	{13, 5, 12, 6, 12, 5},
	{3, 44, 38, 5, 47, 15, 36, 26, 27, 2, 46, 4, 19, 50, 48},
}

func TestSelectionSort(t *testing.T) {
	for _, s := range ss {
		if err := SelectionSort(s); err != nil {
			t.Fatal(err)
		}
		t.Log(s)
	}
}

func TestBubbleSort(t *testing.T) {
	for _, s := range ss {
		if err := BubbleSort(s); err != nil {
			t.Fatal(err)
		}
		t.Log(s)
	}
}

func TestInsertionSort(t *testing.T) {
	for _, s := range ss {
		if err := InsertionSort(s); err != nil {
			t.Fatal(err)
		}
		t.Log(s)
	}
}
