package main

import (
	"fmt"
	"sync"
)

func main() {
	var muA, muB sync.Mutex
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		defer wg.Done()
		muA.Lock()
		defer muA.Unlock()
		// A 依赖 B
		muB.Lock()
		defer muB.Unlock()
	}()

	go func() {
		defer wg.Done()
		muB.Lock()
		defer muB.Unlock()
		// A 依赖 B
		muA.Lock()
		defer muA.Unlock()
	}()

	wg.Wait()
	fmt.Println("结束了")
}
