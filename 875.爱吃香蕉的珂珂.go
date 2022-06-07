/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] 爱吃香蕉的珂珂
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
    max := 0
    for _, pile := range piles {
        if pile > max {
            max = pile
        }
    }
    return 1 + sort.Search(max-1, func(speed int) bool {
        speed++
        time := 0
        for _, pile := range piles {
            time += (pile + speed - 1) / speed
        }
        return time <= h
    })
}
// @lc code=end

