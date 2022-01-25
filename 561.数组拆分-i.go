/*
 * @lc app=leetcode.cn id=561 lang=golang
 *
 * [561] 数组拆分 I
 */

// @lc code=start
func arrayPairSum(nums []int) (res int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		res += nums[i]
	}
	return
}

// @lc code=end

