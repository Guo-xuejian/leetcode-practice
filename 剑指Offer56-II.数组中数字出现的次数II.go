// 剑指 Offer 56 - II. 数组中数字出现的次数 II
// 在一个数组 nums 中除一个数字只出现一次之外，其他数字都出现了三次。请找出那个只出现一次的数字。

// 示例 1：

// 输入：nums = [3,4,3,3]
// 输出：4
// 示例 2：

// 输入：nums = [9,1,7,9,7,9,7]
// 输出：1

// 限制：

// 1 <= nums.length <= 10000
// 1 <= nums[i] < 2^31

func singleNumber(nums []int) int {
	res := 0
	for i := 0; i < 32; i++ {
		bit := 0 //记录该位上的和
		for _, num := range nums {
			bit += ((num >> i) & 1)
		}
		res += ((bit % 3) << i)
	}
	return res
}