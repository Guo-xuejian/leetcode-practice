/*
 * @lc app=leetcode.cn id=908 lang=golang
 *
 * [908] 最小差值 I
 */

// @lc code=start
func smallestRangeI(nums []int, k int) int {
    maxNum, minNum := nums[0], nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] > maxNum {
            maxNum = nums[i]
        } else if nums[i] < minNum {
            minNum = nums[i]
        }
    }
    return max(0, maxNum - minNum - 2 * k)
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
// @lc code=end

