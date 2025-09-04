package main

import (
	"fmt"
)

func worker(id int, f func(int) int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- f(j)
	}

}

func multiplay(val int) int {
	return val * 10
}

func main() {

	const numJobbs = 10
	jobs := make(chan int, numJobbs)
	results := make(chan int, numJobbs)

	for i := 0; i < 3; i++ {
		go worker(i, multiplay, jobs, results)
	}

	for j := 0; j < numJobbs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 0; i < numJobbs; i++ {
		fmt.Println(<-results)
	}

}
