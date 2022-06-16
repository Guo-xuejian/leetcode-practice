/*
 * @lc app=leetcode.cn id=532 lang=golang
 *
 * [532] 数组中的 k-diff 数对
 */

// @lc code=start
func findPairs(nums []int, k int) int {
	visited := map[int]struct{}{}
	res := map[int]struct{}{}
	for _, num := range nums {
		if _, ok := visited[num-k]; ok {
			res[num-k] = struct{}{}
		}
		if _, ok := visited[num+k]; ok {
			res[num] = struct{}{}
		}
		visited[num] = struct{}{}
	}
	return len(res)
}

// @lc code=end

