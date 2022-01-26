/*
 * @lc app=leetcode.cn id=2013 lang=golang
 *
 * [2013] 检测正方形
 */

// @lc code=start
type DetectSquares struct {
	pointMap map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{}
}

func (this *DetectSquares) Add(point []int) {
	if _, ok := this.pointMap[point[1]]; ok {
		this.pointMap[point[1]][point[0]]++
	} else {
		this.pointMap[point[1]] = map[int]int{point[0]: 1}
	}
}

func (this *DetectSquares) Count(point []int) (res int) {
	x, y := point[0], point[1]
	if _, ok := this.pointMap[y]; ok == false {
		return
	}
	yCnt := this.pointMap[y]

	for col, colCnt := range this.pointMap {
		if col != y {
			height := col - y
			res += colCnt[x] * yCnt[x+height] * colCnt[x+height] // 右边
			res += colCnt[x] * yCnt[x-height] * colCnt[x-height] // 左边
		}
	}
	return
}

/**
 * Your DetectSquares object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(point);
 * param_2 := obj.Count(point);
 */
// @lc code=end

