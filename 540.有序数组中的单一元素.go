/*
 * @lc app=leetcode.cn id=540 lang=golang
 *
 * [540] 有序数组中的单一元素
 */

// @lc code=start
func singleNonDuplicate(nums []int) (res int) {
	// 简洁位运算
	// for _, num := range nums {
	//     res ^= num
	// }
	// return

	// 二分
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		// 最终的结果实际上是相同数字起点奇偶性的分界点，左边每对数字的起点索引是偶数，右边为奇数
		// 偶数索引假设为左，比较后者元素，相同则意味着结果应该在 mid 之后
		// 技术索引假设为右，比较前者元素，相同则在 mid 之前
		if nums[mid] == nums[mid^1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nums[left]
}

// @lc code=end

