/*
 * @lc app=leetcode.cn id=859 lang=golang
 *
 * [859] 亲密字符串
 */

// @lc code=start
func buddyStrings(s, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		seen := [26]bool{}
		for _, ch := range s {
			if seen[ch-'a'] {
				return true
			}
			seen[ch-'a'] = true
		}
		return false
	}

	first, second := -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}

// @lc code=end

