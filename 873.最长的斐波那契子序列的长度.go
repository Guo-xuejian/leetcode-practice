/*
 * @lc app=leetcode.cn id=873 lang=golang
 *
 * [873] 最长的斐波那契子序列的长度
 */

// @lc code=start
func lenLongestFibSubseq(arr []int) (ans int) {
    n := len(arr)
    indices := make(map[int]int, n)
    for i, x := range arr {
        indices[x] = i
    }
    dp := make([][]int, n)
    for i := range dp {
        dp[i] = make([]int, n)
    }
    for i, x := range arr {
        for j := n - 1; j >= 0 && arr[j]*2 > x; j-- {
            if k, ok := indices[x-arr[j]]; ok {
                dp[j][i] = max(dp[k][j]+1, 3)
                ans = max(ans, dp[j][i])
            }
        }
    }
    return
}

func max(a, b int) int {
    if b > a {
        return b
    }
    return a
}
// @lc code=end

