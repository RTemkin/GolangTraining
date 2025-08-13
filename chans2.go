package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func gochan(ch chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	ch <- n * n
}

func gochan2(ch chan int, wg *sync.WaitGroup, n int) {
	defer wg.Done()
	ch <- n * n * 2
}

func main() {
	var wg sync.WaitGroup

	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	n := 5

	wg.Add(2)
	go gochan(ch1, &wg, n)
	go gochan2(ch2, &wg, n)

	wg.Wait() // Ждать завершения обеих горутин

	select {
	case val := <-ch1:
		fmt.Println("Ch1", val)
	case val := <-ch2:
		fmt.Println("Ch2", val)
	default:
		fmt.Println("No val")
	}

	colors := []string{"red", "blue", "orange", "yellow"}

	// color := colors[rand.Intn(3)]

	switch colors[rand.Intn(3)] {
	case "red":
		fmt.Println("Red")
	case "blue":
		fmt.Println("blue")
	case "orange":
		fmt.Println("orange")
	case "yellow":
		fmt.Println("yellow")
	default:
		fmt.Println("No color")
	}

}
