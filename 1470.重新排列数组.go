/*
 * @lc app=leetcode.cn id=1470 lang=golang
 *
 * [1470] 重新排列数组
 */

// @lc code=start
func shuffle(nums []int, n int) (res []int) {
	for i := 0; i < n; i++ {
		res = append(res, nums[i])
		res = append(res, nums[i+n])
	}
	return
}

// @lc code=end

