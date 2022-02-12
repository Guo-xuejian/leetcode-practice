/*
 * @lc app=leetcode.cn id=1020 lang=golang
 *
 * [1020] 飞地的数量
 */

// @lc code=start
func numEnclaves(grid [][]int) (res int) {
    // 不是飞地有两种情况，一种常规在内部可以连接至边界，一种是就在边界
    // 那么就从边界开始遍历，深度优先搜索，遍历出所有满足条件的陆地格子，最后内部剩下的没有被访问的陆地格子就是飞地
    m, n := len(grid), len(grid[0])
    directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}    // 四个方向
    var dfs func(i, j int)
    dfs = func(i, j int) {
        // 索引越界（边界必定不是飞地），海洋格子，当前陆地格子已经被访问过
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] == 0 || grid[i][j] == 3 {
            return
        }
        grid[i][j] = 3  // 3 代表已经访问
        for _, direction := range directions {  // 遍历四个方向
            dfs(i + direction[0], j + direction[1])
        }
    }
    for i := 0; i < m; i++ {
        dfs(i, 0)       // 每一行的左边界
        dfs(i, n - 1)   // 右边界
    }
    for j := 0; j < n; j++ {
        dfs(0, j)       // 每一列的上边界
        dfs(m - 1, j)   // 下边界
    }
    for i := 1; i < m - 1; i++ {
        for j := 1; j < n - 1; j++ {
            // 深度优先遍历之后仍然为 1 说明该陆地格子为飞地
            if grid[i][j] == 1 {
                res++
            }
        }
    }
    return
}
// @lc code=end

