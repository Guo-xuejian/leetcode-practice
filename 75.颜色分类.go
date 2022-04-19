/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int) {
	// sort.Slice(nums, func(i, j int) bool {
	//     return nums[i] < nums[j]
	// })
	index := swap2Position(nums, 0)
	swap2Position(nums[index:], 1)
}

func swap2Position(nums []int, target int) (count int) {
	for idx, num := range nums {
		if num == target {
			nums[count], nums[idx] = nums[idx], nums[count]
			count++
		}
	}
	return
}

// @lc code=end

