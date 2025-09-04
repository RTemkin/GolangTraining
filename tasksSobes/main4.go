package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	//var wg sync.WaitGroup

	go func() {
		for i := 0; i < 10; i++ {
			naturals <- i
		}
		close(naturals)
	}()

	go func() {
		for sqar := range naturals {
			squares <- sqar * sqar
		}

		close(squares)
	}()

	for resVal := range squares {
		fmt.Println(resVal)
	}

}
