/*
 * @lc app=leetcode.cn id=334 lang=golang
 *
 * [334] 递增的三元子序列
 */

// @lc code=start
func increasingTriplet(nums []int) bool {
	length := len(nums)
	if length < 3 {
		return false
	}
	minSlice := make([]int, len(nums))
	maxSlice := make([]int, len(nums))
	for i, j := 0, length-1; i < length && j >= 0; i, j = i+1, j-1 {
		if i == 0 {
			minSlice[0] = nums[0]
		} else {
			minSlice[i] = min(nums[i], minSlice[i-1])
		}
		if j == length-1 {
			maxSlice[j] = nums[j]
		} else {
			maxSlice[j] = max(nums[j], maxSlice[j+1])
		}
	}
	for j := 1; j < length-1; j++ {
		if nums[j] > minSlice[j] && nums[j] < maxSlice[j] {
			return true
		}
	}

	// 三循环，限制后者大于前者，小于等于皆不可以
	// for i := 0; i < length-2; i++ {
	//     for j := i+1; j < length-1; j++ {
	//         if nums[j] <= nums[i] {
	//             continue
	//         }
	//         for h := j+1; h < length; h++ {
	//             if nums[h] <= nums[j] {
	//                 continue
	//             }
	//             return true
	//         }
	//     }
	// }
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

