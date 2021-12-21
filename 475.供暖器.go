/*
 * @lc app=leetcode.cn id=475 lang=golang
 *
 * [475] 供暖器
 */

// @lc code=start
func findRadius(houses, heaters []int) (ans int) {
	sort.Ints(heaters)
	for _, house := range houses {
		j := sort.SearchInts(heaters, house+1) // 加一靠右选择
		minDis := math.MaxInt32
		if j < len(heaters) {
			minDis = heaters[j] - house
		}
		i := j - 1
		if i >= 0 && house-heaters[i] < minDis {
			minDis = house - heaters[i]
		}
		if minDis > ans { // 取较大值
			ans = minDis
		}
	}
	return
}

// @lc code=end

