package main

import (
	"fmt"
	"sync"
)

func mergeChans(chs ...<-chan int) <-chan int {
	mergeChan := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chs))
	for _, ch := range chs {
		go func(ch <-chan int, wg *sync.WaitGroup) {
			for num := range ch {
				mergeChan <- num
			}
			wg.Done()
		}(ch, &wg)

	}

	go func() {
		wg.Wait()
		close(mergeChan)
	}()

	return mergeChan
}

func main() {

	a := make(chan int)
	b := make(chan int)
	c := make(chan int)

	go func() {
		for _, num := range []int{10, 20, 30} {
			a <- num
		}
		close(a)
	}()

	go func() {
		for _, num := range []int{11, 22, 33} {
			b <- num
		}
		close(b)
	}()

	go func() {
		for _, num := range []int{111, 222, 333} {
			c <- num
		}
		close(c)
	}()

	for data := range mergeChans(a, b, c) {
		fmt.Println(data)
	}
}
