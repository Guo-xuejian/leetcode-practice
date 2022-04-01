/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] 两个数组间的距离值
 */

// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) (res int) {
    numMap := map[int]bool{}
    for _, num := range arr2 {
        numMap[num] = true
    }
    var twoSideCheck func(num int) bool
    twoSideCheck = func(num int) bool {
        for i := 0; i <= d; i++ {
            if numMap[num + i] || numMap[num - i] {
                return false
            }
        }
        return true
    }
    for _, num := range arr1 {
        if twoSideCheck(num) {
            res++
        }
    }
    return
}
// @lc code=end

