/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 在长度 2N 的数组中找出重复 N 次的元素
 */

// @lc code=start
func repeatedNTimes(nums []int) (res int) {
	sort.Ints(nums)
	if nums[len(nums)/2] == nums[len(nums)-1] {
		res = nums[len(nums)/2]
	} else {
		res = nums[len(nums)/2-1]
	}
	return
}

// @lc code=end

