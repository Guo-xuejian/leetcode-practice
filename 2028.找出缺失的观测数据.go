/*
 * @lc app=leetcode.cn id=2028 lang=golang
 *
 * [2028] 找出缺失的观测数据
 */

// @lc code=start
func missingRolls(rolls []int, mean, n int) []int {
    missingSum := mean * (n + len(rolls))
    for _, roll := range rolls {
        missingSum -= roll
    }
    if missingSum < n || missingSum > n*6 {
        return nil
    }

    quotient, remainder := missingSum/n, missingSum%n
    ans := make([]int, n)
    for i := range ans {
        ans[i] = quotient
        if i < remainder {
            ans[i]++
        }
    }
    return ans
}

// @lc code=end

