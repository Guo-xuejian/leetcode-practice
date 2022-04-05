package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// func main() {
// 	wg.Add(2)
// 	fmt.Println("Work Start Now!")
// 	fooChan := make(chan struct{})
// 	barChan := make(chan struct{})
// 	go printFoo(fooChan, barChan)
// 	go printBar(fooChan, barChan)
// 	fooChan <- struct{}{}
// 	fmt.Println("Work Done!")
// }

func printFoo(fooChan, barChan chan struct{}) {
	defer close(fooChan)
	defer wg.Done()
	for i := 0; i < 101; i++ {
		<-fooChan
		fmt.Println("'foo")
		barChan <- struct{}{}
	}
}

func printBar(fooChan, barChan chan struct{}) {
	defer close(barChan)
	defer wg.Done()
	for i := 0; i < 101; i++ {
		<-barChan
		fmt.Println("'bar")
		if i == 100 {
			return
		}
		fooChan <- struct{}{}
	}
}
