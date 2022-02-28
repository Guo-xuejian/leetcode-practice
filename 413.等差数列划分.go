/*
 * @lc app=leetcode.cn id=413 lang=golang
 *
 * [413] 等差数列划分
 */

// @lc code=start
func numberOfArithmeticSlices(nums []int) (res int) {
    if len(nums) < 3 {
        return
    }
    d, t := nums[0] - nums[1], 0
    for i := 2; i < len(nums); i++ {
        if nums[i - 1] - nums[i] == d {
            t++
        } else {
            d = nums[i - 1] - nums[i]
            t = 0
        }
        res += t
    }
    return
}
// @lc code=end

