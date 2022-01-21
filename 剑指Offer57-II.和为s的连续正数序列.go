// 剑指 Offer 57 - II. 和为s的连续正数序列
// 输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。

// 序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。

// 示例 1：

// 输入：target = 9
// 输出：[[2,3,4],[4,5]]
// 示例 2：

// 输入：target = 15
// 输出：[[1,2,3,4,5],[4,5,6],[7,8]]

// 限制：

// 1 <= target <= 10^5

func findContinuousSequence(target int) (res [][]int) {
	left, right, sumNow := 1, 2, 3
	for left < right {
		if sumNow == target {
			curr := []int{}
			for i := left; i < right+1; i++ {
				curr = append(curr, i)
			}
			res = append(res, curr)
		}
		// 大于或者等于都继续，左边界右移，缩小序列
		if sumNow >= target {
			sumNow -= left
			left += 1
		} else { // 小于右边界右移，扩张序列
			right++
			sumNow += right
		}
	}
	return res
}