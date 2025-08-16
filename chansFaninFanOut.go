package main

import (
	"fmt"
	"sync"
)

func GetInputChan() <-chan int {
	input := make(chan int, 100)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	go func() {
		for _, num := range numbers {
			input <- num
		}
		close(input)
	}()

	return input
}

func GetSquareChan(input <-chan int, name string) <-chan int {

	squarsChane := make(chan int, 100)

	go func() {
		for val := range input {
			fmt.Println(name, val)
			squarsChane <- val * val
		}

		close(squarsChane)
	}()

	return squarsChane
}

func merge(outputsChans ...<-chan int) <-chan int {
	mergeChan := make(chan int, 100)

	var wg sync.WaitGroup

	wg.Add(len(outputsChans))

	outfunc := func(ch <-chan int) {
		for out := range ch {
			mergeChan <- out
		}
		wg.Done()
	}

	for _, output := range outputsChans {
		go outfunc(output)
	}

	go func() {
		wg.Wait()
		close(mergeChan)
	}()

	return mergeChan

}

func main() {

	chanInputNums := GetInputChan()

	SquareChan1 := GetSquareChan(chanInputNums, "SquareChan1")
	SquareChan2 := GetSquareChan(chanInputNums, "SquareChan2")
	SquareChan3 := GetSquareChan(chanInputNums, "SquareChan3")

	merges := merge(SquareChan1, SquareChan2, SquareChan3)

	sumSq := 0

	for val := range merges {
		fmt.Println(val)
		sumSq += val

	}

	fmt.Println("SumSq", sumSq)
}
