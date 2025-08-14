package main

import (
	"errors"
	"fmt"
	"time"
)

var start time.Time

func init() {
	start = time.Now()
}

func service1(ch chan string) {
	fmt.Println("started servise1", time.Since(start))
	ch <- "response service 1"
}

func service2(ch chan string) {
	fmt.Println("started servise2", time.Since(start))
	ch <- "response service 2"
}

func main() {
	fmt.Println("started main", time.Since(start))
	ch1 := make(chan string)
	ch2 := make(chan string)
	err := errors.New("Ничего не поступило")

	go service1(ch1)
	go service2(ch2)


	// select {
	// case res := <-ch1:
	// 	fmt.Println(res, time.Since(start))
	// case res := <-ch2:
	// 	fmt.Println(res, time.Since(start))
	// default:
	// 	fmt.Println("default", err, time.Since(start))
	// }

	

	select {
	case res := <- ch1:
		fmt.Println(res, time.Since(start))
	case res := <- ch2:
		fmt.Println(res, time.Since(start))
	case <-time.After(10 *time.Millisecond):
		fmt.Println("default", err, time.Since(start))
	}

	fmt.Println("stop main", time.Since(start))
}
