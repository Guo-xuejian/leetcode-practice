/*
 * @lc app=leetcode.cn id=1346 lang=golang
 *
 * [1346] 检查整数及其两倍数是否存在
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	numMap := map[int]int{}
	for _, num := range arr {
		numMap[num]++
		if num == 0 && numMap[num] >= 2 {
			return true
		} else if num != 0 && (numMap[num*2] >= 1 || (numMap[num/2] >= 1 && (num/2)*2 == num)) {
			return true
		}
	}
	return false
}

// @lc code=end

