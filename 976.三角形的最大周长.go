/*
 * @lc app=leetcode.cn id=976 lang=golang
 *
 * [976] 三角形的最大周长
 */

// @lc code=start
func largestPerimeter(nums []int) (res int) {
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		// 两小边之和大于大边
		if nums[i-2]+nums[i-1] > nums[i] {
			res = nums[i-2] + nums[i-1] + nums[i]
			break
		}
	}
	return
}

// @lc code=end

