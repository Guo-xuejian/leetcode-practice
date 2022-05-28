/*
 * @lc app=leetcode.cn id=468 lang=golang
 *
 * [468] 验证IP地址
 */

// @lc code=start
func validIPAddress(queryIP string) string {
    if sp := strings.Split(queryIP, "."); len(sp) == 4 {
        for _, s := range sp {
            if len(s) > 1 && s[0] == '0' {
                return "Neither"
            }
            if v, err := strconv.Atoi(s); err != nil || v > 255 {
                return "Neither"
            }
        }
        return "IPv4"
    }
    if sp := strings.Split(queryIP, ":"); len(sp) == 8 {
        for _, s := range sp {
            if len(s) > 4 {
                return "Neither"
            }
            if _, err := strconv.ParseUint(s, 16, 64); err != nil {
                return "Neither"
            }
        }
        return "IPv6"
    }
    return "Neither"
}
// @lc code=end

