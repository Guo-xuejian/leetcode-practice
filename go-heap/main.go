package main

import (
	"container/heap"
	"fmt"
)

type hp []int

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(num interface{}) {
	*h = append(*h, num.(int))
}

func (h *hp) Pop() (res interface{}) {
	*h, res = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func main() {
	arrPre := []int{4,2,5,1,3,7,5,3,8,1,2,5,7}
	// 使用 copy 方法直接拷贝
	arr := make(hp, len(arrPre))
	copy(arr, arrPre)
	length := len(arr)
	heap.Init(&arr)
	fmt.Println(arr)
	for i := 0; i < length; i++ {
		fmt.Println((heap.Pop(&arr)).(int))
	}
	fmt.Println("The program is over......")
}