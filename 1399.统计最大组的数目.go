/*
 * @lc app=leetcode.cn id=1399 lang=golang
 *
 * [1399] 统计最大组的数目
 */

// @lc code=start
func countLargestGroup(n int) int {
	numMap := make(map[int][]int) // 记录总和次数
	for i := 1; i < n+1; i++ {
		key := 0
		a := i
		for {
			if a == 0 {
				break
			}
			key += a % 10 // 整百整千都是 10 的倍数
			a /= 10
		}
		numMap[key] = append(numMap[key], i)
	}
	resultMap := make(map[int]int) // 总和值与长度 map
	maxIndex := 0                  // 最大值索引
	for _, numSlice := range numMap {
		if len(numSlice) > maxIndex {
			maxIndex = len(numSlice) // 获取并列最多的组的索引
		}
		resultMap[len(numSlice)] += 1
	}

	return resultMap[maxIndex]
}

// @lc code=end

