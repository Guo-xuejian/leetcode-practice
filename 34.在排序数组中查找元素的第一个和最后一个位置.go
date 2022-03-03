/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	startIndex, endIndex := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == target && startIndex == -1 {
			startIndex, endIndex = i, i
		} else if nums[i] == target {
			endIndex = i
		} else if startIndex > 0 && endIndex > 0 && nums[i] != target {
			break
		}
	}
	return []int{startIndex, endIndex}
}

// @lc code=end

