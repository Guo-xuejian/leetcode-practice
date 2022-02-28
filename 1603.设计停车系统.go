/*
 * @lc app=leetcode.cn id=1603 lang=golang
 *
 * [1603] 设计停车系统
 */

// @lc code=start
type ParkingSystem struct {
	Big        int
	Medium     int
	Small      int
	CurrBig    int
	currMedium int
	currSmall  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	// 其余默认为 0 即可
	return ParkingSystem{
		Big:    big,
		Medium: medium,
		Small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1 && this.CurrBig < this.Big {
		this.CurrBig++
	} else if carType == 2 && this.currMedium < this.Medium {
		this.currMedium++
	} else if carType == 3 && this.currSmall < this.Small {
		this.currSmall++
	} else {
		return false
	}
	return true
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
// @lc code=end

