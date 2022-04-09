/*
 * @lc app=leetcode.cn id=804 lang=golang
 *
 * [804] 唯一摩尔斯密码词
 */

// @lc code=start
var morse = []string{
    ".-", "-...", "-.-.", "-..", ".", "..-.", "--.",
    "....", "..", ".---", "-.-", ".-..", "--", "-.",
    "---", ".--.", "--.-", ".-.", "...", "-", "..-",
    "...-", ".--", "-..-", "-.--", "--..",
}

func uniqueMorseRepresentations(words []string) int {
    set := map[string]struct{}{}
    for _, word := range words {
        trans := &strings.Builder{}
        for _, ch := range word {
            trans.WriteString(morse[ch-'a'])
        }
        set[trans.String()] = struct{}{}
    }
    return len(set)
}

// @lc code=end

