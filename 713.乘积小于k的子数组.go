/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] 乘积小于K的子数组
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	prod, i := 1, 0
	for j, num := range nums {
		prod *= num
		for ; i <= j && prod >= k; i++ {
			prod /= nums[i]
		}
		ans += j - i + 1
	}
	return
}

// @lc code=end

