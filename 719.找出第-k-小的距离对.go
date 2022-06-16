/*
 * @lc app=leetcode.cn id=719 lang=golang
 *
 * [719] 找出第 k 小的距离对
 */

// @lc code=start
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	return sort.Search(nums[len(nums)-1]-nums[0], func(mid int) bool {
		cnt := 0
		for j, num := range nums {
			i := sort.SearchInts(nums[:j], num-mid)
			cnt += j - i
		}
		return cnt >= k
	})
}

// @lc code=end

