/*
 * @lc app=leetcode.cn id=478 lang=golang
 *
 * [478] 在圆内随机生成点
 */

// @lc code=start
type Solution struct {
	radius, xCenter, yCenter float64
}

func Constructor(radius, xCenter, yCenter float64) Solution {
	return Solution{radius, xCenter, yCenter}
}

func (s *Solution) RandPoint() []float64 {
	r := math.Sqrt(rand.Float64())
	sin, cos := math.Sincos(rand.Float64() * 2 * math.Pi)
	return []float64{s.xCenter + r*cos*s.radius, s.yCenter + r*sin*s.radius}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
// @lc code=end

