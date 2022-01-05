/*
 * @lc app=leetcode.cn id=1400 lang=golang
 *
 * [1400] 构造 K 个回文字符串
 */

// @lc code=start
func canConstruct(s string, k int) bool {
	sLength := len(s)
	if sLength < k {
		// 长度小于则必然不可能拆成 k 个字符串
		return false
	} else if sLength == k {
		// 长度相等，则全部拆成单个字符，单个字符也是回文
		return true
	}
	// 字典记录每个字符出现的次数
	letterMap := make(map[byte]int)
	for _, v := range s {
		letterMap[byte(v)]++
	}
	singleLetterNum := 0 // 单数字符个数
	for _, v := range letterMap {
		if v%2 != 0 {
			singleLetterNum++
		}
	}
	if singleLetterNum > k {
		// 如果单数字符大于 k ，则无法通过分配单个字符为回文串或者单个字符与双数字符组合方式得到 k 个字符串
		return false
	}
	return true
}

// @lc code=end

