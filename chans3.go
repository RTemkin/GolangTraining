package main

import (
	"fmt"
	"runtime"
)

func squares(ch chan int, done chan struct{}) {

	for val := range ch {
		// val := <-ch
		fmt.Println("squar:", val*val)
	}
	fmt.Println("squsres stopped")
	done <- struct{}{}

}

func main() {
	ch := make(chan int, 3)
	done := make(chan struct{})

	go squares(ch, done)
	fmt.Println("active goroutines1", runtime.NumGoroutine())

	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5

	fmt.Println("active goroutines2", runtime.NumGoroutine())

	go squares(ch, done)

	fmt.Println("active goroutines3", runtime.NumGoroutine())

	ch <- 6
	ch <- 7
	ch <- 8
	ch <- 9

	fmt.Println("active goroutines4", runtime.NumGoroutine())

	close(ch)
	<-done
	<-done

	fmt.Println("active goroutines5", runtime.NumGoroutine())

	fmt.Println("main stopped")

}
