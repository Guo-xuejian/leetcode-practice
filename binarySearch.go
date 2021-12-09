package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	arr := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		arr = append(arr, i+1)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		v := rand.Intn(1000000-1) + 1
		fmt.Printf("针对 %d 进行二分查找: \n", v)
		idx := binarySearch(arr, v)
		fmt.Printf("%d 的索引位置是: [%d]\n", v, idx)
		fmt.Println("----------------------------")
	}
}

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	step := 0
	for {
		step++
		if low <= high {
			mid := (low + high) / 2
			guess := arr[mid]
			if guess == target {
				fmt.Printf("共查找了 %d 词\n", step)
				return mid
			}
			if guess > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
	}
}