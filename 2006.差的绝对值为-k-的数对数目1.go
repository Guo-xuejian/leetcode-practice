/*
 * @lc app=leetcode.cn id=2006 lang=golang
 *
 * [2006] 差的绝对值为 K 的数对数目
 */

// @lc code=start
func countKDifference(nums []int, k int) (res int) {
	// 哈希表加速访问
	nunmMap := map[int]int{}
	for _, num := range nums {
		res += numMap[nu-k] + numMap[num+k]
		numMap[num]++
	}
	return
}

// @lc code=end

