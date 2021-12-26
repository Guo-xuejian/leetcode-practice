/*
 * @lc app=leetcode.cn id=2012 lang=golang
 *
 * [2012] 数组美丽值求和
 */

// @lc code=start
func sumOfBeauties(nums []int) (res int) {
	length := len(nums)
	maxArray := []int{nums[0]}
	for i := 1; i < length-2; i++ {
		maxArray = append(maxArray, max(maxArray[len(maxArray)-1], nums[i]))
	}
	minArray := []int{nums[length-1]}
	for i := length - 2; i > 1; i-- {
		minArray = append(minArray, min(minArray[len(minArray)-1], nums[i]))
	}
	for i, j := 0, len(minArray)-1; i < j; i, j = i+1, j-1 {
		minArray[i], minArray[j] = minArray[j], minArray[i]
	}

	for i := 1; i < length-1; i++ {
		if maxArray[i-1] < nums[i] && nums[i] < minArray[i-1] {
			res += 2
		} else if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
			res++
		} else {
			res += 0
		}
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

// @lc code=end

