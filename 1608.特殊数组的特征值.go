/*
 * @lc app=leetcode.cn id=1608 lang=golang
 *
 * [1608] 特殊数组的特征值
 */

// @lc code=start
func specialArray(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	m := nums[n-1]
	for x := 0; x <= m; x++ {
		j := 0
		for ; j < n && nums[j] < x; j++ {
		}
		if n-j == x {
			return x
		}
	}
	return -1
}

// @lc code=end

