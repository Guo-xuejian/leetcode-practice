/*
 * @lc app=leetcode.cn id=1791 lang=golang
 *
 * [1791] 找出星型图的中心节点
 */

// @lc code=start
func findCenter(edges [][]int) (res int) {
	m := len(edges) + 1
	tempSlice := make([]int, m)
	for _, edge := range edges {
		tempSlice[edge[0] - 1]++
		tempSlice[edge[1] - 1]++
	}

	for idx, num := range tempSlice {
		if num == m - 1 {
			res = idx + 1
			break
		}
	}
	return
}
// @lc code=end

