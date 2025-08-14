package main

import (
	"fmt"
	"sync"
	"time"
)

var a int

func sqarWorker(tasks <-chan int, results chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range tasks {
		time.Sleep(10 * time.Millisecond)
		results <- val * val
		fmt.Println("tasck", id)
	}
}

func worker(wg *sync.WaitGroup, mu *sync.Mutex){
	mu.Lock()
	a = a+1
	mu.Unlock()

	wg.Done()
}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex

	tasks := make(chan int, 10)
	results := make(chan int, 10)


	for i := 0; i < 3; i++ {
		wg.Add(1)
		go sqarWorker(tasks, results, i, &wg)
	}

	for i := 0; i < 5; i++ {
		val := i * 2
		tasks <- val
	}
	close(tasks)
	wg.Wait()

	for i := 0; i < 5; i++ {
		res := <-results
		fmt.Println(i, res)
	}

	fmt.Println("--------------------")

	for i := 0; i <1000; i++{
		wg.Add(1)
		go worker(&wg, &mu)
	}

	wg.Wait()

	fmt.Println("worker total amount", a)

	fmt.Println("main stoped")
}
