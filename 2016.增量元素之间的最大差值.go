/*
 * @lc app=leetcode.cn id=2016 lang=golang
 *
 * [2016] 增量元素之间的最大差值
 */

// @lc code=start
func maximumDifference(nums []int) (res int) {
	res = -1
	n := len(nums)
	// for i := 0; i < n; i++ {
	//     for j := i + 1; j < n; j++ {
	//         if nums[i] < nums[j] {
	//             if nums[j] - nums[i] > res {
	//                 res = nums[j] - nums[i]
	//             }
	//         }
	//     }
	// }
	// return
	preMin := nums[0]
	for i := 1; i < n; i++ {
		if nums[i] > preMin {
			if nums[i]-preMin > res {
				res = nums[i] - preMin
			}
		} else {
			preMin = nums[i]
		}
	}
	return
}

// @lc code=end

