/*
 * @lc app=leetcode.cn id=398 lang=golang
 *
 * [398] 随机数索引
 */

// @lc code=start
type Solution struct {
	nums []int // 切片模拟
}

// Constructor 构造函数
func Constructor(nums []int) Solution {
	return Solution{nums: nums}
}

func (this *Solution) Pick(target int) int {
	temp := []int{}
	for idx, val := range this.nums {
		if val == target { // 符合条件记录索引
			temp = append(temp, idx)
		}
	}
	// 随机返回
	return temp[rand.Intn(len(temp))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
// @lc code=end

