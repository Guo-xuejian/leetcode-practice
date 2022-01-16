/*
 * @lc app=leetcode.cn id=1985 lang=golang
 *
 * [1985] 找出数组中的第 K 大整数
 */

// @lc code=start
func kthLargestNumber(nums []string, k int) string {
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		// 较大数字在前，Go 不支持负数索引，所以不能像 Python 一样用 [-k] 解决问题
		if len(a) != len(b) {
			return len(a) > len(b)
		}
		return a > b
	})
	// 零位注意， -1
	return nums[k-1]
}

// @lc code=end

