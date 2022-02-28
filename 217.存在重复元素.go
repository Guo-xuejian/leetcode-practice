/*
 * @lc app=leetcode.cn id=217 lang=golang
 *
 * [217] 存在重复元素
 */

func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// @lc code=end

