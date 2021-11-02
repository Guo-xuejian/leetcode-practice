/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	// 从中心出发，从每一个点开始，当作中心，向外扩散，
	if s == "" {
		return s
	}
	// start ==> 最终的左边界  end ==> 有边界（但是注意数组下标需要最终 +1）
	// left right 分别对应单中心和双中心遍历类型
	var start, end, left1, right1, left2, right2 int
	for i := 0; i < len(s); i++ {
		// 有两种方式，一种是单中心，也就是最终的长度为奇数，另外一种是双中心，也就是最终的长度为偶数，都需要
		left1, right1 = expandAroundCenter(s, i, i)   // 单中心
		left2, right2 = expandAroundCenter(s, i, i+1) // 双中心
		// 判定长度，如果合适即需要重置
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	// 根据条件截取字符串返回
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	// 先决条件	左指针不能越过右指针
	// 判定条件 符合回文条件则继续，左指针向左扩张，右指针向右扩张
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	// 这里 +1 是为了消除 for 循环的递增条件的影响
	return left + 1, right - 1
}

// @lc code=end

