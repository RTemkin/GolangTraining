package main

import "fmt"

func produser(id int, value *int, f func(*int) int, outChan chan<- int) {
	outChan <- f(value)
	fmt.Println("Отправлено Produser id", id)
}

func consumer(id int, inChan <-chan int, resChan chan<- int) {
	for value := range inChan {
		resChan <- value * value
		fmt.Println("Отправлен result id", id)
	}
}

func forProd(val *int) int {
	*val = *val + 10
	return *val
}

func main() {

	const quantGoroot = 3
	tasks := 10
	outChan := make(chan int, tasks)
	resChan := make(chan int, tasks)
	val := 0

	for i := 0; i < tasks; i++ {
		produser(i, &val, forProd, outChan)
	}

	close(outChan)

	for i := 1; i <= 3; i++ {
		go consumer(i, outChan, resChan)
	}

	for i := 0; i < tasks; i++ {
		fmt.Println(<-resChan)
	}

	close(resChan)
}
