/*
 * @lc app=leetcode.cn id=1719 lang=golang
 *
 * [1719] 重构一棵树的方案数
 */

// @lc code=start
func checkWays(pairs [][]int) int {
    adj := map[int]map[int]bool{}
    for _, p := range pairs {
        x, y := p[0], p[1]
        if adj[x] == nil {
            adj[x] = map[int]bool{}
        }
        adj[x][y] = true
        if adj[y] == nil {
            adj[y] = map[int]bool{}
        }
        adj[y][x] = true
    }

    // 检测是否存在根节点
    root := -1
    for node, neighbours := range adj {
        if len(neighbours) == len(adj)-1 {
            root = node
            break
        }
    }
    if root == -1 {
        return 0
    }

    ans := 1
    for node, neighbours := range adj {
        if node == root {
            continue
        }

        currDegree := len(neighbours)
        parent := -1
        parentDegree := math.MaxInt32
        // 根据 degree 的大小找到 node 的父节点 parent
        for neighbour := range neighbours {
            if len(adj[neighbour]) < parentDegree && len(adj[neighbour]) >= currDegree {
                parent = neighbour
                parentDegree = len(adj[neighbour])
            }
        }
        if parent == -1 {
            return 0
        }
        // 检测 neighbours 是否为 adj[parent] 的子集
        for neighbour := range neighbours {
            if neighbour != parent && !adj[parent][neighbour] {
                return 0
            }
        }

        if parentDegree == currDegree {
            ans = 2
        }
    }
    return ans
}
// @lc code=end

