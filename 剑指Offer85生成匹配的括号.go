// 正整数 n 代表生成括号的对数，请设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

// 示例 1：
// 输入：n = 3
// 输出：["((()))","(()())","(())()","()(())","()()()"]

// 示例 2：
// 输入：n = 1
// 输出：["()"]

// 提示：
// 1 <= n <= 8

func generateParenthesis(n int) (res []string) {
	if n <= 0 {
		return
	}
	dfs("", 0, 0, n, &res)
	return
}

// 只要有两种情况的遍历，那就是树
func dfs(path string, openNum, closeNum, limit int, res *[]string) {
	// 左括号优先出现，但不超过 n，右括号有效则不会超过左括号数量
	if openNum > limit || closeNum > openNum {
		return
	}

	if len(path) == 2*limit {
		*res = append(*res, path)
		return
	}

	dfs(path+"(", openNum+1, closeNum, limit, res)
	dfs(path+")", openNum, closeNum+1, limit, res)
}