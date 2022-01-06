/*
 * @lc app=leetcode.cn id=963 lang=golang
 *
 * [963] 最小面积矩形 II
 */

// @lc code=start
// func minAreaFreeRect(points [][]int) float64 {
// 	hash := map[[2]int]struct{}{}
// 	for _, p := range points {
// 		hash[[...]int{p[0], p[1]}] = struct{}{}
// 	}
// 	n := len(points)
// 	ans := 0.0
// 	for i := 0; i < n; i++ {
// 		for j := i + 1; j < n; j++ {
// 			for k := 0; k < n; k++ {
// 				if k == i || k == j {
// 					continue
// 				}
// 				if area, ok := getArea(points[i], points[j], points[k], hash); ok {
// 					if ans == 0.0 || ans > area {
// 						ans = area
// 					}
// 				}

// 			}
// 		}

// 	}
// 	return ans
// }

// func getArea(x, y, z []int, hash map[[2]int]struct{}) (float64, bool) {
// 	a := (x[1] - y[1]) * (y[1] - z[1])
// 	b := (x[0] - y[0]) * (y[0] - z[0])
// 	//a==-b 说明xy线段和yz 线段垂直,不垂直就不是矩形j
// 	if a != -b {
// 		return 0.0, false
// 	}
// 	//寻找例外一个点
// 	qx := x[0] + z[0] - y[0]
// 	qy := x[1] + z[1] - y[1]
// 	//如果不存在例外的点直接返回
// 	if _, ok := hash[[...]int{qx, qy}]; !ok {
// 		return 0.0, false
// 	}
// 	//斜边长度计算相乘就是面积
// 	area := math.Hypot(float64(x[1]-y[1]), float64(x[0]-y[0])) * math.Hypot(float64(y[1]-z[1]), float64(y[0]-z[0]))
// 	return area, true
// }
func minAreaFreeRect(points [][]int) float64 {
	hash := map[[2]int]struct{}{}
	for _, point := range points {
		hash[[2]int{point[0], point[1]}] = struct{}{}
	}
	n := len(points)
	res := math.MaxFloat64
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := 0; k < n; k++ {
				if k == i || k == j { // 点不能重合
					continue
				}
				// 根据矩形对角点之和相等计算出对应的第四个点坐标
				currNeed := [2]int{points[j][0] + points[k][0] - points[i][0], points[j][1] + points[k][1] - points[i][1]}
				if _, ok := hash[currNeed]; ok { // 如果第四个点在字典中存在
					// 对应的点存在
					vx0 := points[j][0] - points[i][0]
					vx1 := points[j][1] - points[i][1]
					vy0 := points[k][0] - points[i][0]
					vy1 := points[k][1] - points[i][1]

					if vx0*vy0+vx1*vy1 == 0 { // 根据向量积判定是否是一个真正的矩形
						// 向量积为0
						// 计算出一个点的两条邻边，矩形的话乘起来就好
						currArea := math.Sqrt(float64(vx0*vx0+vx1*vx1)) * math.Sqrt(float64(vy0*vy0+vy1*vy1))
						if currArea < res {
							// fmt.Println(points[i],points[j],points[k])
							// fmt.Println("width:", float64(vx0*vx0 + vx1 * vx1))
							// fmt.Println("height:", float64(vy0 * vy0 + vy1*vy1))
							res = currArea
						}
					}
				}
			}
		}
	}
	if res == math.MaxFloat64 { // 未出现最小值则返回 0
		return 0.0
	} else {
		return res // 出现了则返回最小值
	}
}

// @lc code=end

