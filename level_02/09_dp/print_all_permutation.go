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
