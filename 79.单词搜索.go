/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
func exist(board [][]byte, word string) bool {
	len1 := len(board)
	len2 := len(board[0])
	len3 := len(word)
	var dfs func(i, j, k int) bool
	dfs = func(i, j, k int) (res bool) {
		if !((i >= 0) && (i < len1)) || !((j >= 0) && (j < len2)) || board[i][j] != word[k] {
			return false
		}
		// 当前字符符合要求
		if k == len3-1 {
			return true
		}
		board[i][j] = byte(1) // 为提高效率，避免后续的操作重复，这一部分剪枝
		res = dfs(i-1, j, k+1) || dfs(i+1, j, k+1) || dfs(i, j-1, k+1) || dfs(i, j+1, k+1)
		board[i][j] = word[k] // 这一个起点结束不能影响下一个全新的起点，所以需要恢复这个值
		return
	}
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if dfs(i, j, 0) { // 遍历所有节点
				return true
			}
		}
	}
	return false
}

// @lc code=end

