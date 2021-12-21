/*
 * @lc app=leetcode.cn id=888 lang=golang
 *
 * [888] 公平的糖果交换
 */

// @lc code=start
func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	sumAlice := 0
	setAlice := map[int]struct{}{}
	for _, v := range aliceSizes {
		sumAlice += v
		setAlice[v] = struct{}{}
	}
	sumBob := 0
	for _, v := range bobSizes {
		sumBob += v
	}
	delta := (sumAlice - sumBob) / 2
	// 差值，也就是较小一方中整体差 delta
	// 只要有一个数满足加上delta在较大方找到
	// 那就直接交换对应数字即可，就可以满足双方相等
	for i := 0; ; i++ {
		y := bobSizes[i]
		x := y + delta
		if _, has := setAlice[x]; has {
			return []int{x, y}
		}
	}
}

// @lc code=end

