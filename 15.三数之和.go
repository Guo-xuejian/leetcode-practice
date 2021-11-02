/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 三数之和
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	numOfArray := len(nums)
	sort.Ints(nums)
	result := [][]int{}

	// i ==> 第一个数字索引
	// j ==> 第二个数字索引
	// h ==> 第三个数字索引
	for i := 0; i < numOfArray; i++ {
		// 此次遍历应该与上一次的数字不同，否则就是无效的
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// h 指向最右端
		h := numOfArray - 1
		target := -nums[i]

		// i+1 在 i 的右侧遍历，避免重复
		for j := i + 1; j < numOfArray; j++ {
			// 同样不能对重复值做判定（因为排过序，所以会快速排除相同数字）
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// 持续缩小范围，结束条件为索引 j 和 h 对应的数字之和小于等于 target 或者索引出现重复，同时重复情况需要在下方做出处理
			for j < h && nums[j]+nums[h] > target {
				h--
			}
			// 如果索引重复，则退出
			if j == h {
				break
			}
			// 不重复且两数之和满足条件，将该种情况写入 result
			if nums[j]+nums[h] == target {
				result = append(result, []int{nums[i], nums[j], nums[h]})
			}
		}
	}
	return result
}

// @lc code=end

