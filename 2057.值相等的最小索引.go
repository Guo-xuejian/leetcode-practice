/*
 * @lc app=leetcode.cn id=2057 lang=golang
 *
 * [2057] 值相等的最小索引
 */

// @lc code=start
func smallestEqual(nums []int) int {
	for index, num := range nums {
		if index%10 == num {
			return index
		}
	}
	return -1
}

// @lc code=end

