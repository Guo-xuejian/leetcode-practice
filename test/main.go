package main

import "fmt"

// func main() {
// 	a := [3]int{1, 2, 3}
// 	for k, v := range a {
// 		if k == 0 {
// 			a[0], a[1] = 100, 200
// 			fmt.Println(a)
// 		}
// 		a[k] = 100 + v
// 	}
// 	fmt.Println(a)
// 	b := []int{1, 2, 3}
// 	for k, v := range b {
// 		if k == 0 {
// 			b[0], b[1] = 100, 200
// 			fmt.Println(b)
// 		}
// 		b[k] = v + 100
// 	}
// 	fmt.Println(b)
// 	c := []int{1, 2, 3}
// 	for k, v := range c {
// 		c = append(c, k+v)
// 	}
// 	fmt.Println(c)
// }

// var DogChan = make(chan struct{}, 1)
// var CatChan = make(chan struct{}, 1)
// var FishChan = make(chan struct{}, 1)
// var wg sync.WaitGroup

// func main() {
// 	fmt.Println("start printing")
// 	wg.Add(3)
// 	DogChan <- struct{}{}
// 	go DogPrint()
// 	go CatPrint()
// 	go FishPrint()
// 	wg.Wait()
// 	fmt.Println("printing finish")
// }

// func DogPrint() {
// 	for i := 0; i < 100; i++ {
// 		<-DogChan
// 		fmt.Println("dog", i)
// 		CatChan <- struct{}{}
// 	}
// 	wg.Done()
// }

// func CatPrint() {
// 	for i := 0; i < 100; i++ {
// 		<-CatChan
// 		fmt.Println("cat", i)
// 		FishChan <- struct{}{}
// 	}
// 	wg.Done()
// }

// func FishPrint() {
// 	for i := 0; i < 100; i++ {
// 		<-FishChan
// 		fmt.Println("fish", i)
// 		DogChan <- struct{}{}
// 	}
// 	wg.Done()
// }

// func test(x interface{}) {
// 	if x == nil {
// 		fmt.Println("nil interface")
// 		return
// 	}
// 	fmt.Println("non-nil interface")
// }

// func main() {
// 	var x *int = nil
// 	test(x)
// }

// func main() {
// 	testMap := make(map[bool]int)
// 	testSlice := make([]int, 1)
// 	mapModfiy(testMap)
// 	fmt.Printf("%p\n", testMap)
// 	sliceModify(testSlice)
// 	fmt.Printf("%p\n", testSlice)
// 	fmt.Println(testMap, &testSlice)
// }

// func mapModfiy(testMap map[bool]int) {
// 	testMap[true] = 1
// 	fmt.Printf("%p\n", testMap)
// }

// func sliceModify(testSlice []int) {
// 	testSlice[0] = 100
// 	fmt.Printf("%p\n", testSlice)
// }

func main() {
	arr1 := [3]int{1, 2, 3}
	arr2 := arr1
	arr2[0] = 100
	fmt.Println(arr1, arr2)

	arr3 := []int{1, 2, 3, 4}
	arr4 := arr3
	arr4 = append(arr4, 89) // 扩容了，指向发生改变
	arr4[0] = 100
	fmt.Println(arr3, arr4)

	map1 := map[int]int{1: 1}
	map2 := map1
	map2[2] = 2
	fmt.Println(map1, map2)
}
