/*
 * @lc app=leetcode.cn id=717 lang=golang
 *
 * [717] 1比特与2比特字符
 */

// @lc code=start
func isOneBitCharacter(bits []int) bool {
    index, m := 0, len(bits)
    for index < m - 1 {
        index += bits[index] + 1
    }
    return index == m - 1
}
// @lc code=end

