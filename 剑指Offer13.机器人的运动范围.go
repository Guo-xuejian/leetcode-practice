// 剑指 Offer 13. 机器人的运动范围
// 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

//  

// 示例 1：

// 输入：m = 2, n = 3, k = 1
// 输出：3
// 示例 2：

// 输入：m = 3, n = 1, k = 0
// 输出：1
// 提示：

// 1 <= n,m <= 100
// 0 <= k <= 20


func movingCount(m int, n int, k int) int {
    q := [][2]int{{0, 0}}
    visited := map[[2]int]bool{{0, 0}: true}
    directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
    res := 1
    isValid := func(x, y int) bool {
        if !(0 <= x && x < m && 0 <= y && y < n) {
            return false
        }
        judgeNum := 0
        for x != 0 {
            judgeNum += x % 10
            x /= 10
        }
        for y != 0 {
            judgeNum += y % 10
            y /= 10
        }
        if judgeNum > k {
            return false
        }
        return true
    }
    for len(q) > 0 {
        length := len(q)
        for i := 0; i < length; i++ {
            point := q[0]
            q = q[1:]
            for _, direction := range directions {
                x, y := point[0] + direction[0], point[1] + direction[1]
                if isValid(x, y) && !visited[[2]int{x, y}] {
                    res++
                    visited[[2]int{x, y}] = true
                    q = append(q, [2]int{x, y})
                }
            }
        }
    }
    return res
}