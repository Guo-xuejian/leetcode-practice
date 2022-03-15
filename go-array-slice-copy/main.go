package main

import (
	"fmt"
	"reflect"
)

func main() {
	arr := [2]int{1, 2}
	test := arr[1:]
	fmt.Println(test)
	fmt.Println(reflect.TypeOf(test))
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
