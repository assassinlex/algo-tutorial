package level_01

import (
	"log"
	"testing"
)

func TestPrintInt32BinaryString(t *testing.T) {
	data := []int32{
		//1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
		//12345, 234344, 555555, 664564,
		//1234567890,
		-2,
	}

	for _, num := range data {
		printInt32BinaryString(num)
		log.Println()
	}
}

func TestGetInt32BinaryNegative(t *testing.T) {
	data := []int32{
		1, 2, 3, 4, 5, -1, -2, -3, -4, -5,
		12345, 234344, 555555, 664564,
		-12345, -234344, -555555, -664564,
		1234567890, -1234567890,
	}

	for _, item := range data {
		log.Println(getInt32BinaryNegative(item))
	}
}

func TestGetFactorialNSum(t *testing.T) {
	data := []int64{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	}

	for _, item := range data {
		log.Println(getFactorialNSum(item))
	}
}
