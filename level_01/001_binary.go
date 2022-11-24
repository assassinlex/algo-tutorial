// 1. 二进制相关的知识
// 2. 位运算

package level_01

import "fmt"

// 输出 int32 的二进制字符串
func printInt32BinaryString(num int32) {

	for i := 31; i >= 0; i-- {
		s := 1
		if (num & (1 << int32(i))) == 0 {
			s = 0
		}
		fmt.Print(s)
	}

	fmt.Print("\t")

}

func printInt64BinaryString(num int64) {

	for i := 63; i >= 0; i-- {
		s := 1
		if (num & (1 << int64(i))) == 0 {
			s = 0
		}
		fmt.Print(s)
	}

	fmt.Print("\t")

}

// getInt32BinaryNegative 获取一个整型的相反数
func getInt32BinaryNegative(num int32) int32 {
	return ^num + 1
}

// getFactorialNSum 获取 N 之前<包含 N> 的阶乘之和，如 1! + 2! + 3! + ... + N!
func getFactorialNSum(n int64) int64 {
	res, cur := int64(0), int64(1)

	for i := int64(1); i <= n; i++ {
		cur *= i
		res += cur
	}

	return res
}
