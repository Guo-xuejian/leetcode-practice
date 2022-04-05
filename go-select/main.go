package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	guoChan := make(chan struct{})
	loveChan := make(chan struct{})
	fangChan := make(chan struct{})
	go func() {
		for i := 0; i < 21; i++ {
			<-guoChan
			chan1 <- "guoxuejian"
			loveChan <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 21; i++ {
			<-loveChan
			chan1 <- "love" 
			fangChan <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 21; i++ {
			<-fangChan
			chan1 <- "fanglu"
			guoChan <- struct{}{}
		}
	}()
	go func() {
		for {
			select {
			case mgs1 := <-chan1:
				fmt.Println(mgs1)
			case fin := <-time.After(time.Second * 10):
				fmt.Println("time-max-limit")
				fmt.Println("current time:", fin)
				return
				// default:
				// 	fmt.Println("waiting for the girl!")
				// 	time.Sleep(time.Second * 1)
			}
		}
	}()
	guoChan <- struct{}{}

	time.Sleep(time.Second * 20)
}
