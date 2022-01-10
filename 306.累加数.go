/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	return dfs(num, 0, 0, 0, 0)
}

func dfs(num string, index, count, prevprev, prev int) bool {
	if index >= len(num) {
		return count > 2
	}
	current := 0
	for i := index; i < len(num); i++ {
		c := num[i]
		if num[index] == '0' && i > index {
			return false // 0 不能为前导符
		}
		current = current*10 + int(byte(c)-'0')
		if count >= 2 {
			sum := prevprev + prev
			if current > sum {
				return false // 不符合条件
			}
			if current < sum {
				continue // 小于，还可以再加入新的字符进行对应计算
			}
		}
		if dfs(num, i+1, count+1, prev, current) {
			return true
		}
	}
	// 一直小于，则不满足
	return false
}

// @lc code=end

