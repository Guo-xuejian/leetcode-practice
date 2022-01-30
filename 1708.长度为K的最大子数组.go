/*
 * @lc app=leetcode.cn id=1708 lang=golang
 *
 * [1708] 长度为K的最大子数组
 */

// @lc code=start
func largestSubarray(nums []int, k int) []int {
	border := len(nums) - k + 1
	maxNumIndex := maxIndex(&nums, border)
	return nums[maxNumIndex : maxNumIndex+k]

}

func maxIndex(nums *[]int, border int) (res int) {
	// 1 <= nums[i] <= 109 所以 res 初始值不必担心
	maxNum := 0
	for i := 0; i < border; i++ {
		curr := (*nums)[i]
		if curr > maxNum {
			maxNum = curr
			res = i
		}
	}
	return
}

// @lc code=end

