/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) (res [][]int) {
	res = append(res, []int{}) // 空集
	// LeetCode 题目虽然说是可以任意顺序返回，实际不行，所以先排序
	for _, num := range nums {
		for _, part := range res { // Go 中的 for range 存在语法糖，所以不需要担心添加的问题
			curr := make([]int, len(part))
			copy(curr, part)
			curr = append(curr, num)
			sort.Ints(curr)
			res = append(res, curr)
		}
	}
	return
}

// @lc code=end

