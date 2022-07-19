package _9_dp

/**
 * 机器人问题: 暴力递归改动态规划 [记忆化搜索]
 *
 * 假设有排成一行的 N 个位置, 记为 1 ~ N, N 一定大于等于 2
 * 开始时机器人在其中的 M 位置上[M 一定是 1 ~ N 中的一个]
 * 若机器人来到 1 的位置, 那么下一步只能往右边来到 2 的位置
 * 若机器人来打 N 的位置, 那么下一步只能往左边来到 N-1 的位置
 * 若机器人来到中间位置[除开 1 & N 的位置], 那么下一步可以往左边或者往右边走
 * 规定机器人必须走 K 步, 最终能来到 P 位置[P 一定是 1 ~ N 中的一个]的方法有多少种
 * 给定四个参数 N, M, K, P, 返回方法种数
 */

func Way1(M, K, P, N int) int {
	if N < 2 || M < 1 || M > N || P < 1 || P > N || K < 1 {
		return -1 // 表示无效
	}
	return way1(M, K, P, N)
}

// way1
// 参数说明: M[当前位置]  K[剩余步数]  P[最终位置]  N[位置数目], P & N 参数不变, M & K 参数可变
func way1(M, K, P, N int) int {
	// 最后一步走完
	if K == 0 {
		res := 0
		if M == P { // 若走完最后一步恰好到 P 位置, 是有效方法, 返回 1
			res = 1
		}
		return res
	}
	// 潜台词: 还未走到最后一步
	if M == 1 {
		return way1(2, K-1, P, N) // M 在 1 的位置只能向 2 的方向走
	}
	if M == N {
		return way1(N-1, K-1, P, N) // M 在 N 的位置只能向 N-1 的方向走
	}
	// M 在中间往左走和往右走的方法种数之和
	return way1(M-1, K-1, P, N) + way1(M+1, K-1, P, N)
}

// way1 函数分析:
// N & P 两个参数不变, 可以看做是 M & K 两个变量的函数, 记作 func(M, K)
// 假设当前位置 M = 3, 需要步数 K = 4 走向目标点 P
// 那么过程可以分解为往左和往右两个过程 func(2, 3), func(4, 3)
// 以此类推, 往左的左 func(1, 2), 往左的右 func (3, 2), 往右的左 func(3, 2), 往右的右 func(5, 2)
// 我们发现计算过程中出现了相同的计算 func(3, 2)
// 暴力递归就是将所有出现的情形都计算出结果, 那么过程中出现的重复计算, 我们可以将其结果缓存起来
// 相同的计算过程直接从缓存中获取结果, 这就是动态规划的一种

func Way2(M, K, P, N int) int {
	if N < 2 || M < 1 || M > N || P < 1 || P > N || K < 1 {
		return -1
	}
	dp := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		dp[i] = make([]int, K+1)
		for j := 0; j <= K; j++ {
			dp[i][j] = -1 // 标识未经理该种计算, 结果未保存
		}
	}
	return way2(M, K, P, N, dp)
}

// way2
// 参数说明: dp[每次计算的结果]
func way2(M, K, P, N int, dp [][]int) int {
	if dp[M][K] != -1 {
		return dp[M][K]
	}
	if K == 0 {
		res := 0
		if M == P {
			res = 1
		}
		return res
	}
	var ans int
	if M == 1 {
		ans = way1(2, K-1, P, N)
	}
	if M == N {
		ans = way1(N-1, K-1, P, N)
	} else {
		ans = way1(M-1, K-1, P, N) + way1(M+1, K-1, P, N)
	}
	dp[M][K] = ans
	return ans
}
