/*
 * @lc app=leetcode.cn id=1995 lang=golang
 *
 * [1995] 统计特殊四元组
 */

// @lc code=start
func countQuadruplets(nums []int) (res int) {
	numsLength := len(nums)
	for a := 0; a < numsLength; a++ {
		for b := a + 1; b < numsLength; b++ {
			for c := b + 1; c < numsLength; c++ {
				for d := c + 1; d < numsLength; d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						res++
					}
				}
			}
		}
	}
	return
}

// @lc code=end

