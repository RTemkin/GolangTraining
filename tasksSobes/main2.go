package main

import (
	"fmt"
	"math/rand"
)

func randNumsGenerator(n int) <-chan int {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make(chan int)

	go func() {
		for i := 0; i < n; i++ {
			r := rand.Intn(1000000)
			out <- r
			// if r < 100000 {
			// 	break
			// }
		}
		close(out)
	}()

	return out

}

func main() {
	i := 0
	for randNum := range randNumsGenerator(10) {
		fmt.Println(randNum, i)
		i++
	}

}
