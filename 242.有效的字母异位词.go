/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
    letterMap := map[byte]int{}
    index := 0
    m, n := len(s), len(t)
    for index < min(m, n) {
        letterMap[s[index]]++
        letterMap[t[index]]--
        index++
    }
    for index < m {
        letterMap[s[index]]++
        index++
    }
    for index < n {
        letterMap[t[index]]--
        index++
    }
    for _, val := range letterMap {
        if val != 0 {
            return false
        }
    }
    return true
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}
// @lc code=end

