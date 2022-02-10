/*
 * @lc app=leetcode.cn id=1404 lang=golang
 *
 * [1404] 将二进制表示减到 1 的步骤数
 */

// @lc code=start
func numSteps(s string) (res int) {
    index := len(s) - 1
    oneStatus := false
    for index > 0 {
        // 当前位为 1 或者当前为应该从 0 变成 1
        if s[index] == '1' || oneStatus {
            res += 2    // 奇数加 1 除 2，两步
            index--
            for index >= 0 && s[index] == '1' {
                // 连接着的所有 1 都变成 0，偶数情况除 2，只需一步
                res++
                index--
            }
            if index >= 0 {
                // 退出上一个循环时当前 index 对应为 0，需要变成 1，oneStatus 记录
                oneStatus = true
            }
        } else {
            // 为 0 ，偶数
            res++
            index--
        }
    }
    return
}
// @lc code=end

