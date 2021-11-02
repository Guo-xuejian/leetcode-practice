/*
 * @lc app=leetcode.cn id=26 lang=golang
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) (slow int) {
	lengthOfNums := len(nums)
	if lengthOfNums == 0 {
		return 0
	}
	// 因为数组本身有序，那么只要该值和上一项相同就删除
	// 快慢指针很好使
	fast := 1
	slow++ // 为了节省空间，函数声明处已经定义了返回值为慢指针
	for fast < lengthOfNums {
		if nums[fast] != nums[fast-1] { // 不相等就覆盖慢指针处元素
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow

}

// @lc code=end

