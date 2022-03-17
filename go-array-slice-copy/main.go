// package main

// import (
// 	"fmt"
// 	"reflect"
// )

// func main() {
// 	testChan := make(chan struct{})
// 	defer close(testChan)
// 	arr := [2]int{1, 2}
// 	test := arr[1:]
// 	test[0] = 200
// 	fmt.Println(test, arr)
// 	fmt.Println(reflect.TypeOf(test))
// 	test1 := make([]int, 4)
// 	copy(test1, arr[:])
// 	fmt.Println("test1", test1)
// 	src := []int{1, 2, 3, 4, 5}
// 	dst := make([]int, 6)

// 	numberOfElementsCopied := copy(dst, src)
// 	fmt.Printf("Number Of Elements Copied: %d\n", numberOfElementsCopied)
// 	fmt.Printf("dst: %v\n", dst)
// 	fmt.Printf("src: %v\n", src)

// 	//After changing numbers2
// 	dst[0] = 10
// 	fmt.Println("\nAfter changing dst")
// 	fmt.Printf("dst: %v\n", dst)
// 	fmt.Printf("src: %v\n", src)
// }
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	catChan := make(chan struct{})
	defer close(catChan)
	dogChan := make(chan struct{})
	defer close(dogChan)
	fishChan := make(chan struct{})
	defer close(fishChan)
	go printCat(catChan, dogChan)
	go printDog(dogChan, fishChan)
	go printFish(fishChan, catChan)
	catChan <- struct{}{}
	wg.Wait()
	fmt.Println("finally over!")
}

func printCat(catChan, dogChan chan struct{}) {
	for i := 0; i < 101; i++ {
		<-catChan
		fmt.Println("cat", i)
		dogChan <- struct{}{}
	}
	wg.Done()
}

func printDog(dogChan, fishChan chan struct{}) {
	for i := 0; i < 101; i++ {
		<-dogChan
		fmt.Println("dog", i)
		fishChan <- struct{}{}
	}
	wg.Done()
}

func printFish(fishChan, catChan chan struct{}) {
	for i := 0; i < 101; i++ {
		<-fishChan
		fmt.Println("fish", i)
		if i != 100 {
			catChan <- struct{}{}
		}
	}
	wg.Done()
}
