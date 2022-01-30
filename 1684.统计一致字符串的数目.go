/*
 * @lc app=leetcode.cn id=1684 lang=golang
 *
 * [1684] 统计一致字符串的数目
 */

// @lc code=start
func countConsistentStrings(allowed string, words []string) (res int) {
	allowedLetterMap := map[byte]int{}
	for i := 0; i < len(allowed); i++ {
		allowedLetterMap[allowed[i]] = 1 // 哈希表加快访问时间，空间换时间
	}
	for _, word := range words {
		controlStatus := true
		for j := 0; j < len(word); j++ {
			if _, ok := allowedLetterMap[word[j]]; !ok { // 有一个不在就不符合条件，退出
				controlStatus = false
				break
			}
		}
		if controlStatus { // 不存在不一致字符，结果 +1
			res++
		}
	}
	return
}

// @lc code=end

