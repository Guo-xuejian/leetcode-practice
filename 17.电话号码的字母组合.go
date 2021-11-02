/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
var num_letter_map map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

// 最终返回的字符串切片
var combinations []string

func letterCombinations(digits string) []string {
	// 为空时需要特别处理，返回空，而不是 ""， 会报错
	if len(digits) == 0 {
		return []string{}
	}
	combinations = []string{}
	backtrack(digits, 0, "")
	return combinations
}

func backtrack(digits string, index int, combination string) {
	if index == len(digits) {
		combinations = append(combinations, combination)
	} else {
		// 通过索引获取字符串某字符时需要处理格式
		digit := string(digits[index])
		letters := num_letter_map[digit]
		lettersCount := len(letters)
		for i := 0; i < lettersCount; i++ {
			// 递归调用，同时 combination+string(letters[i]) 作为一个临时变量传入，记录当前字符组合情况，并在索引值达到边界时写入 combinations
			backtrack(digits, index+1, combination+string(letters[i]))
		}
	}
}

// @lc code=end

