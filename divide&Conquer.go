package main

// 递归 Recursion
// 设立基线条件 Base Case
// 确定 Recursive Case
// 符合基线条件退出递归
// 符合递归条件继续递归

// 步骤
// 1. 找到一个简单的基线条件（Base Case）
// 2. 把问题分开处理，直到它变成基线条件

// func main() {
// 	total := sum([]int{1, 2, 3, 4, 5, 6, 7})
// 	fmt.Println("数组加总的结果=====>")
// 	fmt.Println(total)
// }

func sum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + sum(arr[1:])
}
