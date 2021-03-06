// 剑指 Offer 03. 数组中重复的数字
// 找出数组中重复的数字。

// 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

// 示例 1：

// 输入：
// [2, 3, 1, 0, 2, 5, 3]
// 输出：2 或 3

// 限制：

// 2 <= n <= 100000

func findRepeatNumber(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return nums[i]
		}
	}
	return -1 // 无所谓，只是需要一个返回值保证程序编译通过
}

// 2022-04-15
// func findRepeatNumber(nums []int) int {
//     numMap := map[int]struct{}{}
//     for _, num := range nums {
//         if _, ok := numMap[num]; ok {
//             return num
//         } else {
//             numMap[num] = struct{}{}
//         }
//     }
//     return -1
// }