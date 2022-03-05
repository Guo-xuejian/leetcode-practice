/*
 * @lc app=leetcode.cn id=2100 lang=golang
 *
 * [2100] 适合打劫银行的日子
 */

// @lc code=start
func goodDaysToRobBank(security []int, time int) (res []int) {
	m := len(security)
	if time == 0 {
		for i := 0; i < m; i++ {
			res = append(res, i)
		}
		return
	}
	for i := time; i < m-time; i++ {
		currMax, currMin := security[i], security[i]
		fitAble := true
		for j := 1; j <= time; j++ {
			if security[i-j] < currMax || security[i+j] < currMin {
				fitAble = false
				break
			}
			currMax, currMin = security[i-j], security[i+j]
		}
		if fitAble {
			res = append(res, i)
		}
	}
	return
}

// @lc code=end

