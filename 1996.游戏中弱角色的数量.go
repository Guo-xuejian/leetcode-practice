/*
 * @lc app=leetcode.cn id=1996 lang=golang
 *
 * [1996] 游戏中弱角色的数量
 */

// @lc code=start
func numberOfWeakCharacters(properties [][]int) (res int) {
    sort.Slice(properties, func(i, j int) bool {
        if properties[i][0] == properties[j][0] {
            return properties[i][1] < properties[j][1]  // 防御值从小到大
        } else {
            return properties[i][0] > properties[j][0]  // 攻击力从大到小
        }
    })
    maxDenfense := 0
    for _, propert := range properties {
        if maxDenfense > propert[1] {   // 比之前的攻击力小的同时防御力也小，当前角色为弱角色
            res++
        } else {
            maxDenfense = max(maxDenfense, propert[1])
        }
    }
    return
}


func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}

// @lc code=end

