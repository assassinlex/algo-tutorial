// 1, 二分法
// 1, 复杂度
// 1, 动态数组
// 1, 哈希表

package level_01

import "math/rand"

// bisection01 有序数组中找到目标值 n
func bisection01(s []int32, n int32) bool {
	length := len(s)
	if s == nil || length == 0 {
		return false
	}
	left, mid, right := 0, 0, length-1

	for left <= right {
		// todo:: 此处会有坑
		mid = (left + right) / 2
		if s[mid] == n {
			return true
		} else if s[mid] < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

// bisectionGen 生成器
func bisectionGen() []int32 {
	panic("implements me...")
}

// bisectionCheck 对数器
func bisectionCheck(s []int32, n int32) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}

	return false
}


// 返回邻值不等的一个数组
func randomArray(maxLength, maxValue int) []int {
	length := rand.Intn(maxLength) + 1
	ret := make([]int, length, length)

	ret[0] = rand.Intn(maxValue)
	for i := 1; i < length; i++ {
		for {
			ret[i] = rand.Intn(maxValue)
			if ret[i] != ret[i-1] { break }
		}
	}

	return ret
}
