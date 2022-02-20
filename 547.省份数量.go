/*
 * @lc app=leetcode.cn id=547 lang=golang
 *
 * [547] 省份数量
 */

// @lc code=start
func findCircleNum(isConnected [][]int) (res int) {
	m := len(isConnected)
	visited := map[int]bool{}
	// 一个城市只会属于一个省份
	var dfs func(i int)
	dfs = func(i int) {
		for j := 0; j < m; j++ {
			// 如果不相连或者已经出现则无需继续
			if isConnected[i][j] == 1 && !visited[j] {
				visited[j] = true // 相连且未出现的新城市
				dfs(j)
			}
		}
	}
	for i := 0; i < m; i++ {
		if !visited[i] { // 新的城市出现
			dfs(i)
			res++
		}
	}
	return
}

// @lc code=end

