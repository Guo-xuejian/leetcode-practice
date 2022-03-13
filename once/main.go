package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var once sync.Once
	onceBody := func() {
		fmt.Println("Only once")
	}
	for i := 0; i < 10; i++ {
		go func() {
			once.Do(onceBody)
		}()
	}
	time.Sleep(time.Second * 4)
}
