package _9_dp

/**
 * 打印字符串全部排列
 */

func permutation01(str string) []string {
	var ans []string
	if len(str) == 0 {
		return ans
	}
	return f1([]byte(str), 0, ans)
}

// f1
// 问题: 获取一个字符串的全部排列
//
// 思路:
//     迭代字符串, 每一个字符都可能出现在原字符串的任何位置
//     固定好一个字符的位置后, 重复固定剩余字符可能出现在的所有剩余位置
//                        [abc]
//           [a..]        [b..]        [c..]        每个字符来到 0 的位置
//        [ab.] [ac.]  [ba.] [bc.]  [ca.] [cb.]     剩下的字符中每个字符来到 1 的位置
//        [abc] [acb]  [bac] [bca]  [cab] [cba]     剩下的字符中每个字符来到 2 的位置
//
// 参数说明:
//     chars: 用户参数
//     i: 字符都有机会出现在 i 的位置
//     ans: 所有排列的集合
func f1(chars []byte, i int, ans []string) []string {
	length := len(chars)
	if i == length { // 当处理到最后一种情况后的 chars, 也是结果的一种
		ans = append(ans, string(chars))
		return ans
	}
	for j := i; j < length; j++ { // 剩余字符从 i 开始, 所有的字符都有机会来到 i 的位置
		chars[j], chars[i] = chars[i], chars[j] // 当 j 来到 i 时
		ans = f1(chars, i+1, ans)               // 处理 i 下一个字符[i+1] 的后续所有可能
		chars[j], chars[i] = chars[i], chars[j] // 还原现场
	}
	return ans
}

func permutation02(str string) []string {
	var ans []string
	if len(str) == 0 {
		return ans
	}
	return f2([]byte(str), 0, ans)
}

// f2
// 问题: 获取一个字符串的全部排列, 且不能出现重复排列 [分支限界]
//
// 注意: 此处虽然是用 set 进行去重, 但是与子序列的去重稍显不同,
//      子序列是将所有的支路都处理完了之后, 将结果存入 set, 利用 set 的特性自动去重
//      分支限界是某种情况发生后, 后序的一系列相关操作都不进行,
//      在这个例子中表达为 i 位置出现了某个字符后, 后续的这个字符再次出现在这个位置上的后序可能结果一概不要
func f2(chars []byte, i int, ans []string) []string {
	length := len(chars)
	if i == length {
		ans = append(ans, string(chars))
	}
	alloc := make([]bool, 256) // 构造字符的分配记录, 利用字符对应的 ascii 在数组中的 bool 值作为依据
	for j := i; j < length; j++ {
		if !alloc[chars[j]] { // 只有 j 位置的字符未被分配过, 才进行排列组合
			alloc[chars[j]] = true
			chars[j], chars[i] = chars[i], chars[j]
			ans = f2(chars, i+1, ans)
			chars[j], chars[i] = chars[i], chars[j]
		}
	}
	return ans
}
