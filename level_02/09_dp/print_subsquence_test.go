package _9_dp

import (
	"fmt"
	"strconv"
	"testing"
)

func TestSubs(t *testing.T) {
	strings := []string{
		"abc",
		"235",
		"adf",
		"aaa",
	}

	for _, str := range strings {
		fmt.Println(Subs(str))
	}
}

func TestSubsNoRepeat(t *testing.T) {
	strings := []string{
		"abc",
		"235",
		"adf",
		"aaa",
	}

	for _, str := range strings {
		s := SubsNoRepeat(str)
		for v, _ := range s {
			fmt.Printf("%v ", v)
		}
		fmt.Println()
	}
}

// 切片测试一
func TestSlice01(t *testing.T) {

	strings := make([]string, 0, 5)

	func(s []string) {
		for i := 0; i < 5; i++ {
			s = append(s, strconv.Itoa(i))
			fmt.Println(s)
		}
	}(strings)

	fmt.Println(strings)
}

// 切片测试二
func TestSlice02(t *testing.T) {

	ints := make([]int, 5)

	func(s []int) {
		for i := 0; i < 5; i++ {
			s[i] = i + 1
			fmt.Println(s)
		}
	}(ints)

	fmt.Println(ints)
}

// 切片测试三
func TestSlice03(t *testing.T) {

	ints := make([]int, 5)

	func(s []int) {
		for i := 0; i < 5; i++ {
			s = append(s, i)
			fmt.Println(s)
		}
	}(ints)

	fmt.Println(ints)
}
