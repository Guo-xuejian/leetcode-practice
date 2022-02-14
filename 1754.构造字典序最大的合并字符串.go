/*
 * @lc app=leetcode.cn id=1754 lang=golang
 *
 * [1754] 构造字典序最大的合并字符串
 */

// @lc code=start
func largestMerge(word1 string, word2 string) string {
	m, n := len(word1), len(word2)
	allSlice := make([]byte, m+n)
	idx, idx1, idx2 := 0, 0, 0
	for ; idx1 < m && idx2 < n; idx++ {
		// 对比两个字符串的字典序
		if greaterDictionaryString(word1[idx1:], word2[idx2:]) {
			// a 先出现字典序较大字符，或者 len(a) > len(b) 同时两者至 len(b) - 1 索引字符均相同
			allSlice[idx] = word1[idx1]
			idx1++
		} else {
			// b 先出现字典序较大字符
			allSlice[idx] = word2[idx2]
			idx2++
		}
	}
	// 结束之后可能有一个没有全部写入，分别判定写入，只会有一个真正写入
	for ; idx1 < m; idx1, idx = idx1+1, idx+1 {
		allSlice[idx] = word1[idx1]
	}
	for ; idx2 < n; idx2, idx = idx2+1, idx+1 {
		allSlice[idx] = word2[idx2]
	}
	return string(allSlice)
}

// 是否优先从word1取值
func greaterDictionaryString(a, b string) bool {
	for i := range a {
		// 持续比较的前提是 i 索引之前的元素都一致
		// a 的长度大于 b，直接返回 true，从 a 中取字符
		// 因为之前部分全部相等，那么只有 a 更有可能在后续的字符中拆出字典序较大的字符
		// 同理，a[i] > b[i] a 早一步出现字典序较大的字符
		if i+1 > len(b) || a[i] > b[i] {
			return true
		} else if a[i] < b[i] {
			// a 字典序小于 b，b 先出现字典序更大的字符
			return false
		}
	}
	// a 和 b 完全相等，或者满足这个条件下 len(a) < len(b)
	// 同理，b 更有可能在后续出现字典序较大字符
	return false
}

// @lc code=end

