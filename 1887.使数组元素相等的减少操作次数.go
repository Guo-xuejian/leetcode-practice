/*
 * @lc app=leetcode.cn id=1887 lang=golang
 *
 * [1887] 使数组元素相等的减少操作次数
 */

// @lc code=start
func reductionOperations(nums []int) int {
	sort.Ints(nums)
	length := len(nums)
	// res 总操作次数，curr 每个元素操作次数
	res, curr := 0, 0
	for i := 1; i < length; i++ {
		if nums[i] != nums[i-1] { // 不重复则需要进行数字变换
			curr++
		}
		res += curr // 由于是渐进式修改，所有每次都需要加上新修改的次数
	}
	return res
}

// @lc code=end

