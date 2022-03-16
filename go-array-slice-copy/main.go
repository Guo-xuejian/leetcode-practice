package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [2]int{1, 2}
	test := arr[1:]
	test[0] = 200
	fmt.Println(test, arr)
	fmt.Println(reflect.TypeOf(test))
	test1 := make([]int, 4)
	copy(test1, arr[:])
	fmt.Println("test1", test1)
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, 6)

	numberOfElementsCopied := copy(dst, src)
	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)

	//After changing numbers2
	dst[0] = 10
	fmt.Println("\nAfter changing dst")
	fmt.Printf("dst: %v\n", dst)
	fmt.Printf("src: %v\n", src)
}
