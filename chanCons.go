package main

import "fmt"

func produser(ch chan<- int, n int) {

	for i := 0; i < n; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(ch <-chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int)
	go produser(ch, 10)
	consumer(ch)

}
