/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {
	oneLength := len(nums)
	allRanges := factorial(oneLength)
	res := make([][]int, allRanges)
	index := 0
	for i := 0; i < allRanges; i++ {
		res[i] = make([]int, oneLength)
	}

	var backTrack func(tempNums, path []int, length int) 
	backTrack = func(tempNums, path []int, length int) {
		if length == 0 {
			for i := 0; i < oneLength; i++ {
				res[index][i] = path[i]
			}
			index++
			return
		}

		for i := 0; i < length; i++ {
			curr := tempNums[i]
			path = append(path, curr)
			tempNums = append(tempNums[:i], tempNums[i+1:]...)
			backTrack(tempNums, path, len(tempNums))
			tempNums = append(tempNums[:i], append([]int{curr}, tempNums[i:]...)...)
			path = path[:len(path) - 1]
		}
	}
	backTrack(nums, []int{}, oneLength)
	return res
}

func factorial(num int) int {
	res := 1
	for num != 0 {
		res *= num
		num--
	}
	return res
}
// @lc code=end

