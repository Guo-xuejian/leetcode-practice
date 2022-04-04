/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// @lc code=start
func twoSum(numbers []int, target int) (res []int) {
	numMap := map[int]int{}
	for idx1, num := range numbers {
		if idx2, ok := numMap[target-num]; ok {
			res = []int{idx2 + 1, idx1 + 1}
			break
		}
		numMap[num] = idx1
	}
	return
}

// @lc code=end

