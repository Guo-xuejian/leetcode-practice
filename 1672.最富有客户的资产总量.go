/*
 * @lc app=leetcode.cn id=1672 lang=golang
 *
 * [1672] 最富有客户的资产总量
 */

// @lc code=start
func maximumWealth(accounts [][]int) (res int) {
	for _, account := range accounts {
		res = max(res, sum(account))
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func sum(nums []int) (res int) {
	for _, num := range nums {
		res += num
	}
	return
}

// @lc code=end

