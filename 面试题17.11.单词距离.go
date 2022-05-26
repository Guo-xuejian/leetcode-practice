// 面试题 17.11. 单词距离
// 有个内含单词的超大文本文件，给定任意两个不同的单词，找出在这个文件中这两个单词的最短距离(相隔单词数)。如果寻找过程在这个文件中会重复多次，而每次寻找的单词不同，你能对此优化吗?

// 示例：

// 输入：words = ["I","am","a","student","from","a","university","in","a","city"], word1 = "a", word2 = "student"
// 输出：1
// 提示：

// words.length <= 100000

func findClosest(words []string, word1, word2 string) int {
	ans := len(words)
	index1, index2 := -1, -1
	for i, word := range words {
		if word == word1 {
			index1 = i
		} else if word == word2 {
			index2 = i
		}
		if index1 >= 0 && index2 >= 0 {
			ans = min(ans, abs(index1-index2))
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
