/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] 独特的电子邮件地址
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	eamilMap := map[string]struct{}{}
	for _, email := range emails {
		index := strings.IndexByte(email, '@')
		localName := strings.SplitN(email[:index], "+", 2)[0]
		localName = strings.ReplaceAll(localName, ".", "")
		eamilMap[localName+email[index:]] = struct{}{}
	}
	return len(eamilMap)
}

// @lc code=end

