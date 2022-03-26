/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) (res [][]int) {
	sort.Ints(candidates)
	m := len(candidates)
	var backTrack func(path []int, curr, index int)
	backTrack = func(path []int, curr, index int) {
		if curr < 0 {
			return
		}
		if curr == 0 {
			temp := make([]int, len(path))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		for index < m {
			path = append(path, candidates[index])
			backTrack(path, curr-candidates[index], index)
			path = path[:len(path)-1]
			index++
		}
	}
	backTrack([]int(nil), target, 0)
	return
}

// @lc code=end

