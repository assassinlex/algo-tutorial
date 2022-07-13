package _9_dp

import "math"

/**
 * 四种尝试模型[二]: 范围尝试模型
 */

// 给定一个整形数组 arr, 代表述职不同的纸牌排成一条线
// 玩家 A 和 B 依次拿走每张纸牌
// 规定玩家 A 先拿, 玩家 B 后拿
// 但是每个玩家只能拿走最左或最右的纸牌
// 玩家 A 和 B 都绝顶聪明 [博弈问题, 两人的结果都是很大, 但一定有一个最大]
// 请返回最后获胜者的分数
//
// 说明: 纸牌对所有人可见, 即游戏开始结果就已经定了
//      游戏不是每次都拿左右最大的牌, 而是拿完牌后让对手拿到的牌更小

func win1(s []int) (int, int) {
	l := len(s)
	if s == nil || l == 0 {
		return 0, 0
	}
	first := first1(s, 0, l-1)
	next := next1(s, 0, l-1)
	return first, next
}

// first1 先手: 依次拿到自己的牌[大] + 对手留给自己下一次的牌[小]
func first1(s []int, l, r int) int {
	if l == r { // 最后一张牌
		return s[l]
	}
	p1 := s[l] + next1(s, l+1, r)                  // 先手拿左边的牌和下一次拿牌的结果 [拿完牌后变作后手]
	p2 := s[r] + next1(s, l, r-1)                  // 先手拿右边的牌和下一次拿牌的结果
	return int(math.Max(float64(p1), float64(p2))) // 取大 [让自己拿到的牌最大]
}

// next1 后手: 拿到对手留给自己下一次的牌[小]
func next1(s []int, l, r int) int {
	if l == r { // 只有牌的适量是基数时, l 和 r 才会相等, 最后一张牌肯定是先手的
		return 0
	}
	p1 := first1(s, l+1, r)                        // 后手拿左边的牌 [轮到后手的时候, 后手变作先手]
	p2 := first1(s, l, r-1)                        // 后手拿右边的牌
	return int(math.Min(float64(p1), float64(p2))) // 取小 [让后手拿到的牌最小]
}
