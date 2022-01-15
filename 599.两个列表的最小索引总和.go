/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	list1Length, list2Length := len(list1), len(list2)
	minIndexSum := list1Length + list2Length	// 去一个较大值
	indexSumStringSliceMap := map[int][]string{}
	for i := 0; i < list1Length; i++ {
		for j := 0; j < list2Length; j++ {
			// 当前索引和
			currIndexSum := i + j
			if currIndexSum > minIndexSum {
				// 当前大于，之后只会更大，及时退出，同时大于，即使满足条件也不会是最终的结果
				break
			}
			if list1[i] == list2[j] {
				indexSumStringSliceMap[currIndexSum] = append(indexSumStringSliceMap[currIndexSum], list1[i])
				// 比较获取最小索引和
				if currIndexSum < minIndexSum {
					minIndexSum = currIndexSum
				}
			}
		}
	}
	return indexSumStringSliceMap[minIndexSum]
}

// @lc code=end

