/*
 * @lc app=leetcode.cn id=397 lang=golang
 *
 * [397] 整数替换
 */

// @lc code=start
func integerReplacement(n int) int {
	if n == 1 { // 1 满足条件直接返回
		return 0
	}

	// 偶数则加一（除法的一步）递归
	if n%2 == 0 {
		return 1 + integerReplacement(n/2)
	}

	// 奇数后必定为偶数，所以直接加 2（加法和除法的模拟）继续递归
	return 2 + min(integerReplacement(n/2), integerReplacement(n/2+1))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// @lc code=end

