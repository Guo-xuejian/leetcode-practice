/*
 * @lc app=leetcode.cn id=2023 lang=golang
 *
 * [2023] 连接后等于目标字符串的字符串对
 */

// @lc code=start
func numOfPairs(nums []string, target string) (ans int) {
	cnt := map[string]int{}
	for _, s := range nums {
		cnt[s]++
	}
	for i, n := 1, len(target); i < n; i++ {
		p, s := target[:i], target[i:] // 枚举所有前缀+后缀
		if p != s {
			ans += cnt[p] * cnt[s]
		} else {
			ans += cnt[p] * (cnt[p] - 1) // 前后缀相同时，相当于从 cnt[p] 个下标中选择两个不同下标的排列数，即 A(cnt[p], 2)
		}
	}
	return
}

// @lc code=end

