package main

import "fmt"

// 快速排序 O(nlog2n),比选择排序要快 O(n2)
// 在生活中经常使用
// 使用了 D & C 策略（分而治之）
// divide and Conquer

func quickSort(arr []int) []int {
	// 无值、单值已到达要求
	if len(arr) < 2 {
		return arr
	}

	// 默认以首元素为基准，同时定义左右数组(切片实现)
	pivot := arr[0]
	var left, right []int

	for _, ele := range arr[1:] { // 遍历切分原数组
		if ele <= pivot {
			left = append(left, ele)
		} else {
			right = append(right, ele)
		}
	}
	// 递归调用
	return append(quickSort(left), append([]int{pivot}, quickSort(right)...)...)
}

func main() {
	testArr := []int{6, 4, 7, 9, 3, 1, 4, 5, 3, 2}
	testArr = quickSort(testArr)
	fmt.Println("数组排序之后为====>")
	fmt.Println(testArr)
}
