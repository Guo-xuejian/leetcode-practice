/*
 * @lc app=leetcode.cn id=1356 lang=golang
 *
 * [1356] 根据数字二进制下 1 的数目排序
 */

// @lc code=start
func sortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		count1, count2 := countOnes(arr[i]), countOnes(arr[j])
		if count1 < count2 {
			return true
		} else if count1 == count2 {
			return arr[i] < arr[j]
		}
		return false
	})
	return arr
}

func countOnes(num int) (res int) {
	for num > 0 {
		res += num & 1
		num >>= 1
	}
	return
}
// @lc code=end

