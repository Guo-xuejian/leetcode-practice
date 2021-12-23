// 给定两个整数数组a和b，计算具有最小差绝对值的一对数值（每个数组中取一个值），并返回该对数值的差

// 示例：
// 输入：{1, 3, 15, 11, 2}, {23, 127, 235, 19, 8}
// 输出：3，即数值对(11, 8)

func smallestDifference(a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)
	i, j, diff := 0, 0, 0
	res := math.MaxInt32
	// 排序后方便减少计算量
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			diff = b[j] - a[i]
			i++
		} else {
			diff = a[i] - b[j]
			j++
		}
		if diff < res {
			res = diff
		}
	}
	return res
}

// func quickSort(numx *[]int) (res []int) {
//     left, right := []int{}, []int{}
//     first := (*numx)[0]
//     for i := 0; i<len(*numx); i++ {
//         if (*numx)[i] < first {
//             left = append(left, (*numx)[i])
//         } else {
//             right = append(right, (*numx)[i])
//         }
//     }
//     return append(quickSort(*left), append([]int{first}, quickSort(right)...)...)
// }