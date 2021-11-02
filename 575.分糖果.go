/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candyType []int) (res int) {
	set := map[int]int{} // 记录糖果的种类数量
	for _, t := range candyType {
		set[t] = 0
	}
	res = len(candyType) / 2
	if len(set) < res { // 糖果种类少于可吃的数量则返回种类数量
		res = len(set)
	}
	return
}

// @lc code=end

