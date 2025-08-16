package main

import "fmt"

func fib(length int) <-chan int {

	ch := make(chan int)

	go func() {
		for i, j, l := 0, 1, 0; l < length; i, j, l = i+j, i, l+1 {
			ch <- i
		}
		close(ch)
	}()

	return ch
}

func main() {

	c := fib(10)
	for f := range c {    // или прям так: for f := range fib(10)
		fmt.Println(f)
	}

}
