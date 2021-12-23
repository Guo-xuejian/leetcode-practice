// 实现一个算法，确定一个字符串 s 的所有字符是否全都不同。

// 示例 1：
// 输入: s = "leetcode"
// 输出: false

// 示例 2：
// 输入: s = "abc"
// 输出: true

func isUnique(astr string) bool {
	for i := 0; i < len(astr); i++ {
		// 正向查找和逆向结果不一致则证明存在重复字符
		if strings.Index(astr, string(astr[i])) != strings.LastIndex(astr, string(astr[i])) {
			return false
		}
	}
	return true
}