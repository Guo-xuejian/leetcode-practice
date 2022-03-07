package main

import (
	"fmt"
	"sync"
)

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

var DogChan = make(chan struct{}, 1)
var CatChan = make(chan struct{}, 1)
var FishChan = make(chan struct{}, 1)
var wg sync.WaitGroup

func main() {
	fmt.Println("start printing")
	wg.Add(3)
	DogChan <- struct{}{}
	go DogPrint()
	go CatPrint()
	go FishPrint()
	wg.Wait()
	fmt.Println("printing finish")
}

func DogPrint() {
	for i := 0; i < 100; i++ {
		<-DogChan
		fmt.Println("dog", i)
		CatChan <- struct{}{}
	}
	wg.Done()
}

func CatPrint() {
	for i := 0; i < 100; i++ {
		<-CatChan
		fmt.Println("cat", i)
		FishChan <- struct{}{}
	}
	wg.Done()
}

func FishPrint() {
	for i := 0; i < 100; i++ {
		<-FishChan
		fmt.Println("fish", i)
		DogChan <- struct{}{}
	}
	wg.Done()
}
