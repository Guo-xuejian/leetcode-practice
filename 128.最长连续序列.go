/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) (res int) {
	numMap := map[int]bool{}
	for _, num := range nums {
		numMap[num] = true
	}
	for _, num := range nums {
		if !numMap[num-1] {
			currMax := 1
			for temp := num + 1; numMap[temp]; temp++ {
				currMax++
			}
			if currMax > res {
				res = currMax
			}
		}
	}
	return
}
// @lc code=end

