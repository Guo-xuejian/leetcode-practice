/*
 * @lc app=leetcode.cn id=241 lang=golang
 *
 * [241] 为运算表达式设计优先级
 */

// @lc code=start
const addition, subtraction, multiplication = -1, -2, -3

func diffWaysToCompute(expression string) []int {
	ops := []int{}
	for i, n := 0, len(expression); i < n; {
		if unicode.IsDigit(rune(expression[i])) {
			x := 0
			for ; i < n && unicode.IsDigit(rune(expression[i])); i++ {
				x = x*10 + int(expression[i]-'0')
			}
			ops = append(ops, x)
		} else {
			if expression[i] == '+' {
				ops = append(ops, addition)
			} else if expression[i] == '-' {
				ops = append(ops, subtraction)
			} else {
				ops = append(ops, multiplication)
			}
			i++
		}
	}

	n := len(ops)
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, n)
	}
	var dfs func(int, int) []int
	dfs = func(l, r int) []int {
		res := dp[l][r]
		if res != nil {
			return res
		}
		if l == r {
			dp[l][r] = []int{ops[l]}
			return dp[l][r]
		}
		for i := l; i < r; i += 2 {
			left := dfs(l, i)
			right := dfs(i+2, r)
			for _, x := range left {
				for _, y := range right {
					if ops[i+1] == addition {
						dp[l][r] = append(dp[l][r], x+y)
					} else if ops[i+1] == subtraction {
						dp[l][r] = append(dp[l][r], x-y)
					} else {
						dp[l][r] = append(dp[l][r], x*y)
					}
				}
			}
		}
		return dp[l][r]
	}
	return dfs(0, n-1)
}

// @lc code=end

