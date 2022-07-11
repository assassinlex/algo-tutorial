package _9_dp

import (
	set "algo/level_02/03_map"
	"fmt"
)

/**
 * 打印字符串子序列问题
 */

func Subs(str string) []string {
	return process01([]byte(str), 0, []string{}, "")
}

// process01
// 问题: 获取一个字符串的全部子序列
//
// 思路:
//     迭代字符串, 将 i 位置的字符放入子序列中的结果和不放入子序列产生两种结果
//     将每个位置的决定产生的两种结果全收集
//
// 参数说明:
//     chars: 用户参数
//     i: 即将要处理的字符所在索引[包含两个信息, 0 ~ i-1 的位置已经做完决定, 剩下 i ~ last 继续做决定 ]
//     ans: 所有可能的子序列
//     path: i 处做决定产生的结果继续影响 i + 1 的结果
func process01(chars []byte, i int, ans []string, path string) []string {
	if i == len(chars) { // 已经处理完全, 将当前决定产生的结果放入结果结合
		ans = append(ans, path)
		return ans
	}
	ans = process01(chars, i+1, ans, path)                                // i 处的字符不放入子序
	ans = process01(chars, i+1, ans, fmt.Sprintf("%s%c", path, chars[i])) // i 处的字符放入子序
	return ans
}

func SubsNoRepeat(str string) set.MySet {
	return process02([]byte(str), 0, set.NewSet(0), "")
}

// process02
// 问题: 获取一个字符串的全部子序列, 且不能出现重复子序列
func process02(chars []byte, i int, ans set.MySet, path string) set.MySet {
	if i == len(chars) { // 已经处理完全, 将当前决定产生的结果放入结果结合
		ans.Add(path)
		return ans
	}
	ans = process02(chars, i+1, ans, path)                                // i 处的字符不放入子序
	ans = process02(chars, i+1, ans, fmt.Sprintf("%s%c", path, chars[i])) // i 处的字符放入子序
	return ans
}
