/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	originalLength := len(s)
	uniqeList := getUnqieString(s)
	uniqeListLength := len(uniqeList)
	if originalLength == uniqeListLength {
		return originalLength
	}
	temp := []byte{}
	lengthOfLongestStr := 0
	for i := 0; i < originalLength; i++ {
		if index, ok := isContain(temp, s[i]); ok {
			// 存在， 则计算 temp 中元素数量并和 lengthOfLongestStr 进行比较
			lengthOfLongestStr = maxOne(lengthOfLongestStr, len(temp))
			// 得出 temp 中 s[i] 的索引值并对 temp 进行截断
			temp = temp[index+1:]
			// 还需要加入本次遍历对应元素
			temp = append(temp, s[i])
		} else {
			// 不存在将遍历元素加入 temp
			temp = append(temp, s[i])
		}
	}
	return maxOne(lengthOfLongestStr, len(temp))
}

func getUnqieString(s string) (res string) {
	eleMap := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if _, ok := eleMap[s[i]]; ok {
			continue
		} else {
			eleMap[s[i]] = 0
		}
	}
	for _, val := range eleMap {
		res += fmt.Sprintf("%c", val)
	}
	return
}

func maxOne(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func isContain(x []byte, y byte) (index int, ok bool) {
	for j := 0; j < len(x); j++ {
		if x[j] == y {
			ok = true
			index = j
		}
	}
	return
}

// @lc code=end

