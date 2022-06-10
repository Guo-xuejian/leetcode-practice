/*
 * @lc app=leetcode.cn id=730 lang=golang
 *
 * [730] 统计不同回文子序列
 */

// @lc code=start
func countPalindromicSubsequences(s string) (ans int) {
    const mod int = 1e9 + 7
    n := len(s)
    dp := [4][][]int{}
    for i := range dp {
        dp[i] = make([][]int, n)
        for j := range dp[i] {
            dp[i][j] = make([]int, n)
        }
    }
    for i, c := range s {
        dp[c-'a'][i][i] = 1
    }

    for sz := 2; sz <= n; sz++ {
        for i, j := 0, sz-1; j < n; i++ {
            for k, c := 0, byte('a'); k < 4; k++ {
                if s[i] == c && s[j] == c {
                    dp[k][i][j] = (2 + dp[0][i+1][j-1] + dp[1][i+1][j-1] + dp[2][i+1][j-1] + dp[3][i+1][j-1]) % mod
                } else if s[i] == c {
                    dp[k][i][j] = dp[k][i][j-1]
                } else if s[j] == c {
                    dp[k][i][j] = dp[k][i+1][j]
                } else {
                    dp[k][i][j] = dp[k][i+1][j-1]
                }
                c++
            }
            j++
        }
    }

    for _, d := range dp {
        ans += d[0][n-1]
    }
    return ans % mod
}

// @lc code=end

