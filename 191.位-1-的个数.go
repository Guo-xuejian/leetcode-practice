/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) (res int) {
    for num > 0 {
        if num & 1 == 1 {
            res++
        }
        num >>= 1
    }
    return
}
// @lc code=end

