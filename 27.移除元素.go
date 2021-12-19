/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	left := 0
	for _, v := range nums {
		if val != v {
			nums[left] = v
			left++
		}
	}
	return left
}

// @lc code=end

