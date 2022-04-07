package main

// import (
// 	"container/heap"
// 	"fmt"
// )

// // type hp []int

// // func (h *hp) Len() int {
// // 	return len(*h)
// // }

// // func (h *hp) Less(i, j int) bool {
// // 	return (*h)[i] > (*h)[j]
// // }

// // func (h *hp) Swap(i, j int) {
// // 	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
// // }

// // func (h *hp) Push(x interface{}) {
// // 	*h = append(*h, x.(int))
// // }

// // func (h *hp) Pop() (v interface{}) {
// // 	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
// // 	return
// // }

// // func (h *hp) Peek() interface{} {
// // 	return (*h)[0]
// // }

// type hp []int

// func (h *hp) Len() int {
// 	return len(*h)
// }

// func (h *hp) Less(i, j int) bool {
// 	return (*h)[i] < (*h)[j]
// }

// func (h *hp) Swap(i, j int) {
// 	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
// }

// func (h *hp) Push(x interface{}) {
// 	*h = append(*h, x.(int))
// }

// func (h *hp) Pop() (v interface{}) {
// 	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
// 	return
// }

// func (h *hp) Peek() interface{} {
// 	return (*h)[0]
// }

// func main() {
// 	arr := &hp{}
// 	heap.Init(arr)
// 	arr.Push(23)
// 	arr.Push(25)
// 	arr.Push(21)
// 	arr.Push(13)
// 	fmt.Println(arr)
// 	heap.Init(arr)
// 	arr.Push(20)
// 	fmt.Println(arr)
// 	arr.Push(17)
// 	fmt.Println(arr)
// 	arr.Push(5)
// 	fmt.Println(arr)
// 	fmt.Println(arr.Peek())
// 	fmt.Println(arr.Pop())
// 	fmt.Println(arr.Pop())
// 	fmt.Println(arr.Pop())
// }

import (
	"container/heap"
	"fmt"
)

type hp []int

func main() {
	h := &hp{0, 1, 2, 1} // 创建slice
	// sort.Sort(*h)
	// fmt.Println(*h)
	heap.Init(h)    // 将数组切片进行堆化
	fmt.Println(*h) // [0 1 3 6 2 5 7 9 8 4] 由Less方法可控制小顶堆
	// fmt.Println(heap.Pop(h)) // 调用pop 0 返回移除的顶部最小元素
	// heap.Push(h, 6)          // 调用push [1 2 3 6 4 5 7 9 8] 添加一个元素进入堆中进行堆化
	fmt.Println("new: ", *h) // [1 2 3 6 4 5 7 9 8 6]
	for len(*h) > 0 {        // 持续推出顶部最小元素
		fmt.Printf("%d \n ", heap.Pop(h))
	}
}

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i] < h[j] }

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Pop() interface{} {
	res := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return res
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}
