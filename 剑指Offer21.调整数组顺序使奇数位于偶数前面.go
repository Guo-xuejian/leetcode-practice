// 剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
// 输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。

// 示例：

// 输入：nums = [1,2,3,4]
// 输出：[1,3,2,4]
// 注：[3,1,2,4] 也是正确的答案之一。

// 提示：

// 0 <= nums.length <= 50000
// 0 <= nums[i] <= 10000

func exchange(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}
	sort.Slice(nums, func(i, j int) bool {
		if nums[i]%2 == 1 && nums[j]%2 != 1 {
			return true
		}
		return false
	})
	return nums
}