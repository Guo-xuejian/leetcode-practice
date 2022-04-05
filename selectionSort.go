package main

// func main() {
// 	arr := []int{6, 4, 7, 9, 3, 1, 4, 5, 3, 2}
// 	arr  = selectionSort(arr)
// 	fmt.Println("排序之后的数组==========>")
// 	fmt.Println(arr)
// }

func findSmallest(arr []int) int {
	smallest := arr[0]
	smallestIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIndex = i
		}
	}
	return smallestIndex
}

func selectionSort(arr []int) []int {
	var result []int
	count := len(arr)
	for i := 0; i < count; i++ {
		smallestIndex := findSmallest(arr)
		result = append(result, arr[smallestIndex])
		arr = append(arr[:smallestIndex], arr[smallestIndex+1:]...)
	}

	return result
}
