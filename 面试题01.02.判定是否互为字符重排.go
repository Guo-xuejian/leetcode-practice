// 面试题 01.02. 判定是否互为字符重排
// 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

// 示例 1：

// 输入: s1 = "abc", s2 = "bca"
// 输出: true 
// 示例 2：

// 输入: s1 = "abc", s2 = "bad"
// 输出: false
// 说明：

// 0 <= len(s1) <= 100
// 0 <= len(s2) <= 100

func CheckPermutation(s1 string, s2 string) bool {
    letterMap := map[byte]int{}
    for _, letter := range s1 {
        letterMap[byte(letter)]++
    }
    for _, letter := range s2 {
        letterMap[byte(letter)]--
    }

    for _, val := range letterMap {
        if val != 0 {
            return false
        }
    }
    return true
}