// 1, 二分法
// 1, 复杂度
// 1, 动态数组
// 1, 哈希表

package level_01

import (
	"log"
	"math/rand"
	"testing"
)

func TestBisection01(t *testing.T) {
	data := [][]int32 {
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 22, 33, 44, 55, 66, 67, 78, 89, 110},
	}

	for _, s := range data {
		idx := int32(rand.Intn(len(s)))
		get := bisection01(s, s[idx])
		check := bisectionCheck(s, s[idx])
		if get != check {
			log.Println(s, s[idx])
			log.Printf("expect: %t    get: %t\n", get, check)
		}
	}
}

// 等概率验证
func TestRandIntN(t *testing.T) {
	// n -> times
	data := map[int]int {
		1: 1000,
		2: 100000,
		3: 10000000,
		4: 100000000,
	}

	for n, times := range data {
		log.Printf("[0, 10) 的随机数小于 %d 的概率是: %f\n", n, equalP(times, n))
	}


}

func TestRandomArray(t *testing.T) {
	// max_length -> max_value
	data := map[int]int {
		3: 100,
		6: 500,
		9: 1000,
	}

	for length, value := range data {
		log.Printf("%v\n", randomArray(length, value))
	}
}