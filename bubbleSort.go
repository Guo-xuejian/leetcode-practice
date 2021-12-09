package main

import (
	"fmt"
)

func bubbleSort(arr [5]int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := 0; j < length-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	// return arr
}

func main() {
	arr := [5]int{11, 33, 6, 5, 1}
	// arr = bubbleSort(arr)
	// fmt.Println(arr)
	// go func() {

	// }()
	for {
		go bubbleSort(arr)
	}
	fmt.Println("The Program Start !")
	// for i := 0; i < 6; i++ {
	// 	go fmt.Println(i)
	// }
	// time.Sleep(100000)
}
