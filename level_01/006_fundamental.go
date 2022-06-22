package level_01

// 位运算实现四则运算

// add 加法原理
func add(a, b int) int {
	// 两数相加 = 两数无进位相加 + 进位数据
	// 无进位相加 = 两数异或(^), 进位数据 = 两数与(&) << 1
	// 当进位数据 = 0 时, 异或的结果就是加法的结果
	res := a
	for b != 0 {
		res = a ^ b
		b = (a & b) << 1
		a = res
	}
	return res
}

// minus 减法原理
func minus(a, b int) int {
	// 两数相减 = 被减数 + 减数的相反数(取反 + 1)
	return add(a, add(^b, 1))
}

// multi 乘法原理
func multi(a, b int) int {
	res := 0
	for b != 0 {
		if (b & 1) != 0 {
			res = add(res, a)
		}
		a <<= 1
		b = int(uint(b) >> 1)
	}
	return res
}

// div 出发原理
func div(a, b int) int {
	return 0
}


