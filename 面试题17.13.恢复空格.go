// 面试题 17.13. 恢复空格
// 哦，不！你不小心把一个长篇文章中的空格、标点都删掉了，并且大写也弄成了小写。像句子"I reset the computer. It still didn’t boot!"已经变成了"iresetthecomputeritstilldidntboot"。在处理标点符号和大小写之前，你得先把它断成词语。当然了，你有一本厚厚的词典dictionary，不过，有些词没在词典里。假设文章用sentence表示，设计一个算法，把文章断开，要求未识别的字符最少，返回未识别的字符数。

// 注意：本题相对原题稍作改动，只需返回未识别的字符数

// 示例：

// 输入：
// dictionary = ["looked","just","like","her","brother"]
// sentence = "jesslookedjustliketimherbrother"
// 输出： 7
// 解释： 断句后为"jess looked just like tim her brother"，共7个未识别字符。
// 提示：

// 0 <= len(sentence) <= 1000
// dictionary中总字符数不超过 150000。
// 你可以认为dictionary和sentence中只包含小写字母。

func respace(dictionary []string, sentence string) int {
	dictMap := map[string]int{}
	for _, word := range dictionary {
		dictMap[word] = 0
	}

	length := len(sentence)
	// dp[i] 是 s[0,...,i-1] 的最少未识别字符数 (特别地，dp[0] = 0)
	dp := []int{0}
	for i := 0; i < length; i++ {
		dp = append(dp, length)
	}
	// dp[j] = min{ dp[i], 如果 s[i:j] 在 dictionary 中; dp[i] + j - i, 如果 s[i:j] 不在 dictionary 中 | i < j }
	// 每个索引都对应一个最少未识别字符数，当前字符串，以及对应索引之前的未识别字符数加总就是当前索引上的最少未识别字符串
	for j := 1; j < length+1; j++ {
		for i := 0; i < j; i++ {
			if _, ok := dictMap[string(sentence[i:j])]; ok {
				dp[j] = min(dp[j], dp[i])
			} else {
				dp[j] = min(dp[j], dp[i]+j-i)
			}
		}
	}
	return dp[length]
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}