/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了100.00%的用户
func shortestToChar(s string, c byte) []int {
    m := len(s)
    var getDistance func(index int) int
    getDistance = func(index int) (res int) {
        if s[index] == c {
            return
        }
        left, right := index, index
        for left >= 0 || right < m {
            if left >= 0 {
                if s[left] == c {
                    res = abs(index - left)
                    return
                }
                left--
            }
            if right < m {
                if s[right] == c {
                    res = abs(index - right)
                    return
                }
                right++
            }
        }
        return
    }
    res := make([]int, m)
    for i := 0; i < m; i++ {
        res[i] = getDistance(i)
    }
    return res
}

func abs(num int) int {
    if num < 0 {
        return -num
    }
    return num
}
// @lc code=end

