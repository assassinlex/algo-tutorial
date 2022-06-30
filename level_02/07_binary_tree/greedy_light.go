package _7_binary_tree

import (
	set "algo/level_02/03_map"
	"errors"
	"fmt"
	"math"
	"math/rand"
)

/**
 * 案例: 路灯问题
 *    一个由 X 和 . 组成的字符串代表一条街道,
 *    X 代表的是一堵特殊墙, 无法放置灯, 也无法被照亮,
 *    . 代表的是居民区, 可以被放置灯, 放了灯就被照亮, 一盏灯能够照亮 3 个对象, 自身和左右相邻距离为 1 的 .(即居民区)。
 *    求任意街道的居民区被全部照亮的最小灯数量。
 */

// 解题思路[画图]:
//     假设 idx 是 X, 那么肯定不放灯 [X??  因为不能], 直接去 idx + 1 上继续做决策
//     假设 idx 是 . , 那么决策依赖 idx + 1 的位置
//     假设 idx + 1 是 X, 那么 idx 必须放灯 [.X?  因为灯只能照亮自身、左右相邻距离为 1 的位置], 然后去 idx + 2 上继续做决策
//     假设 idx + 1 是 ., 那么决策依赖 idx + 2 的位置
//     假设 idx + 2 是 X, 那么 idx + 1 必须放灯 [..X idx 也可以放灯, 但不是最优解]
//     假设 idx + 2 是 ., 那么 idx + 1 必须放灯 [... idx + 1 放灯是最优解]
//     其实到这里就可以看出, 无论 idx + 2 上是 . 还是 X, 都不影响结果, 都是在 idx + 1 上放灯, 然后去 idx + 3 上继续做决策

// GetMinLightCnt01 方案一: 暴力模式
func GetMinLightCnt01(road string) int {
	if len(road) == 0 {
		return 0
	}
	return getMinLightCnt([]byte(road), 0, set.NewSet())
}

// 参数说明
//     chars[idx:], idx 所处所有位置自由选择放灯与否
//     chars[0:idx-1], 这中间的位置哪些位置放灯的结果已经有了, 放灯的索引已经存入了 lights 中
//     idx 当前字符索引
//     lights 放了灯的索引集合 [求 lights 最小值]
func getMinLightCnt(chars []byte, idx int, lights set.MySet) int {
	if idx == len(chars) { // 作业结束, 此时灯的放置结果已经决策完成
		for i, v := range chars {
			// 不能放灯的情形
			if v == 'X' {
				continue
			}
			// 可以放灯的情形 [那么放灯的位置必然是 i - 1 || i || i + 1 其中一个]
			if !lights.Contain(i-1) && !lights.Contain(i) && !lights.Contain(i+1) {
				return math.MaxInt32 // 若以上 3 个位置都没有放灯, 返回 int64 最大值表示无效解
			}
		}
		return lights.Size()
	} else { // 作业还未结束, 此时需继续对 chars[idx:] 做决策
		no := getMinLightCnt(chars, idx+1, lights) // idx 不放灯的情形 [chars[idx] 是 X, 或者 是 . 但是选择不放灯]
		yes := math.MaxInt32                       // idx 默认放灯的情形, 以最大值表示
		if chars[idx] == '.' {                     // 放灯的情形
			lights.Add(idx)
			yes = getMinLightCnt(chars, idx+1, lights) // 当确定 idx 放灯时
			lights.Remove(idx)
		}
		return int(math.Min(float64(no), float64(yes)))
	}
}

// GetMinLightCnt02 方法二: 贪心算法
func GetMinLightCnt02(road string) int {
	chars := []byte(road)
	idx := 0
	light := 0

	for idx < len(chars) {
		if chars[idx] == 'X' { // 不能放灯的情形
			idx++
			continue
		}

		light++ // 可以放灯的情形, 能放就 +1, 此时不管是放在 idx 还是 idx + 1 上了, 反正是放了

		if idx+1 == len(chars) { // 此时已经决策完全 [字符串已经到达最后一个]
			break
		} else {
			if chars[idx+1] == 'X' { // 直接去 idx + 2 上继续决策
				idx = idx + 2
			} else { // 直接去 idx + 3 上继续决策
				idx = idx + 3
			}
		}
	}

	return light
}

//func GetMinLightCnt03(road string) int {
//	chars := []byte(road)
//	idx := 0
//	light := 0
//
//	for i, v := range chars {
//
//	}
//}

// 获取随机字符串
func randomString(length int) string {
	chars := make([]byte, rand.Intn(length)+1)
	for i, _ := range chars {
		chars[i] = 'X'
		if rand.Intn(2) == 1 {
			chars[i] = '.'
		}
	}
	return string(chars)
}

// 对数器
func detector(length, times int) error {
	for i := 0; i < times; i++ {
		s := randomString(length)
		ans1 := GetMinLightCnt01(s)
		ans2 := GetMinLightCnt02(s)
		if ans1 != ans2 {
			return errors.New(
				fmt.Sprintf("数据: %s, GetMinLightCnt01: %d, GetMinLightCnt02: %d", s, ans1, ans2))
		}
	}

	return nil
}
