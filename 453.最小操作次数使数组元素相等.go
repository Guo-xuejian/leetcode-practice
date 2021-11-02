/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小操作次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) (res int) {
	min_num := min(&nums) // 计算最小值逐一加总与其他值的差值
	for _, val := range nums {
		res += val - min_num
	}
	return
}

func min(nums *[]int) (res int) { // 最小值函数
	lengthOfNums := len(*nums)
	res = (*nums)[0]
	for i := 1; i < lengthOfNums; i++ {
		if (*nums)[i] < res {
			res = (*nums)[i]
		}
	}
	return
}

// @lc code=end

