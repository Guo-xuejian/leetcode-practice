// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func main() {
// 	count := 0
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	n := 10
// 	wg.Add(n)
// 	for i := 0; i < n; i++ {
// 		go func() {
// 			defer wg.Done()
// 			for j := 0; j < 10000; j++ {
// 				mu.Lock()
// 				count++
// 				mu.Unlock()
// 			}
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println(count)
// 	fmt.Println("计数结束了")
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// func main() {
// 	var m sync.RWMutex
// 	go read(&m, 1)
// 	go read(&m, 2)
// 	go read(&m, 3)
// 	time.Sleep(2 * time.Second)
// }

// func read(m *sync.RWMutex, i int) {
// 	fmt.Println(i, "reader start")
// 	m.RLock()
// 	fmt.Println(i, "reading")
// 	time.Sleep(1 * time.Second)
// 	m.RUnlock()
// 	fmt.Println(i, "reader over")
// }

package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0

func main() {
	var m sync.RWMutex
	for i := 1; i <= 3; i++ {
		go write(&m, i)
	}
	for i := 1; i <= 3; i++ {
		go read(&m, i)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("final count:", count)
}

func read(m *sync.RWMutex, i int) {
	fmt.Println(i, "reader start")
	m.RLock()
	fmt.Println(i, "reading count:", count)
	time.Sleep(1 * time.Millisecond)
	m.RUnlock()
	fmt.Println(i, "reader over")
}

func write(m *sync.RWMutex, i int) {
	fmt.Println(i, "write start")
	m.Lock()
	count++
	fmt.Println(i, "writing")
	time.Sleep(1 * time.Millisecond)
	m.Unlock()
	fmt.Println(i, "write over")
}
