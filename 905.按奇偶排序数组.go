/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity(nums []int) []int {
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] % 2 == 0
    })
    return nums
}
// @lc code=end

