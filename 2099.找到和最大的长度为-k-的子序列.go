/*
 * @lc app=leetcode.cn id=2099 lang=golang
 *
 * [2099] 找到和最大的长度为 K 的子序列
 */

// @lc code=start
func maxSubsequence(nums []int, k int) []int {
	length := len(nums)
	var nums2 []int = make([]int, length)
	copy(nums2, nums)
	sort.Ints(nums2)
	var numFreq map[int]int = make(map[int]int)
	for i := length - 1; i > length-k-1; i-- {
		numFreq[nums2[i]]++
	}
	var res []int = make([]int, 0)
	for _, x := range nums {
		if numFreq[x] > 0 {
			res = append(res, x)
			numFreq[x]--
		}
	}
	return res
}

// @lc code=end

