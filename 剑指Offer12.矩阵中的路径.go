// 剑指 Offer 12. 矩阵中的路径
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。

// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 例如，在下面的 3×4 的矩阵中包含单词 "ABCCED"（单词中的字母已标出）。

// 示例 1：

// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true
// 示例 2：

// 输入：board = [["a","b"],["c","d"]], word = "abcd"
// 输出：false

// 提示：

// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// board 和 word 仅由大小写英文字母组成

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	// 以每一个点作为起点去做匹配，只要存在一个成功即可
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0, word, &board) {
				return true
			}
		}
	}
	return false
}

func dfs(i, j, k int, word string, board *[][]byte) (res bool) {
	// 不能越界
	if i < 0 || j < 0 || i >= len(*board) || j >= len((*board)[0]) || (*board)[i][j] != word[k] {
		return false
	}

	// 字符匹配成功，如果长度已是最大，返回 true
	if k == len(word)-1 {
		return true
	}

	// 剪枝操作，为了以 (*board)[i][j] 为起点的后续匹配创造一个不可能匹配的情况
	(*board)[i][j] = '1' // 题目是 byte
	res = dfs(i+1, j, k+1, word, board) || dfs(i-1, j, k+1, word, board) || dfs(i, j+1, k+1, word, board) || dfs(i, j-1, k+1, word, board)
	(*board)[i][j] = word[k] // 判定结束后恢复该值，不能影响其余起点的判定
	return
}