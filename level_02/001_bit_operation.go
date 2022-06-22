package level_02

import "log"

// swap 交换两个变量的值
func swap(x, y int) (int, int) {
	x = x ^ y
	y = x ^ y
	x = x ^ y

	return x, y
}

// getNum 一段整数序列, 一个数出现奇数次，其他数出现偶数次, 求出现奇数次的数
func getNum(s []int) int {
	eor := 0
	for _, item := range s {
		eor = eor ^ item
	}
	return eor
}

// getAnswer 获取一个数的二进制最右侧的 1 与其右侧的 0 组成的数, 如: 0100 0100 => 0100 => 4
func getAnswer(x int) int {
	return x & ((^x) + 1)
}

func demo() {

	m := []map[string]string{
		{"receiver": "prometheus-alert-center", "status": "firing"},
		{"receiver": "prometheus-alert-center", "status": "firing"},
	}

	for _, item := range m {
		log.Println(item["status"])
	}
}
