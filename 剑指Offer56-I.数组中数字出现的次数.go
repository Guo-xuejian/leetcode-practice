// 剑指 Offer 56 - I. 数组中数字出现的次数
// 一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。

// 示例 1：

// 输入：nums = [4,1,4,6]
// 输出：[1,6] 或 [6,1]
// 示例 2：

// 输入：nums = [1,2,10,4,1,4,3,3]
// 输出：[2,10] 或 [10,2]

// 限制：

// 2 <= nums.length <= 10000

func singleNumbers(nums []int) []int {
	var a int
	for i := range nums {
		a ^= nums[i]
	}
	mask := a & (-a)
	res := make([]int, 2)
	for _, v := range nums {
		if (v & mask) == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}