/*
 * @lc app=leetcode.cn id=139 lang=golang
 *
 * [139] 单词拆分
 */

// @lc code=start
func wordBreak(s string, wordDict []string) bool {
    wordDictSet := make(map[string]bool)
    for _, w := range wordDict {    // 转为字典加速访问
        wordDictSet[w] = true
    }
    dp := make([]bool, len(s) + 1)
    dp[0] = true	// 空字符串也是组成部分
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordDictSet[s[j:i]] {   // s[:j] 存在同时 s[j:i] 存在
                dp[i] = true
                break
            }
        }
    }
    return dp[len(s)]
}

// @lc code=end

