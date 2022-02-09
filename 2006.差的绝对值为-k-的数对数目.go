/*
 * @lc app=leetcode.cn id=2006 lang=golang
 *
 * [2006] 差的绝对值为 K 的数对数目
 */

// @lc code=start
func countKDifference(nums []int, k int) (res int) {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if abs(nums[i], nums[j]) == k {
				res++
			}
		}
	}
	return
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
// @lc code=end

