/*
 * @lc app=leetcode.cn id=1748 lang=golang
 *
 * [1748] 唯一元素的和
 */

// @lc code=start
func sumOfUnique(nums []int) (ans int) {
	// 哈希表计数
	cnt := map[int]int{}
	for _, num := range nums {
		cnt[num]++
	}
	for num, c := range cnt {
		if c == 1 {
			ans += num
		}
	}
	return
}

// @lc code=end

