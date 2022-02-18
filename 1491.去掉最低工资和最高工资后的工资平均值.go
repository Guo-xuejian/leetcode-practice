/*
 * @lc app=leetcode.cn id=1491 lang=golang
 *
 * [1491] 去掉最低工资和最高工资后的工资平均值
 */

// @lc code=start
func average(salary []int) float64 {
    minNum, maxNum, total := salary[0], salary[0], 0
    for _, num := range salary {
        if num > maxNum {
            maxNum = num
        } else if num < minNum {
            minNum = num
        }
        total += num
    }
    return float64(total - maxNum - minNum) / float64(len(salary) - 2)
}
// @lc code=end

