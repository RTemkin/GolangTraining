package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
)

func searthSecretVal(ctx context.Context, start, fin, secretVal, idx int, ch chan int, wg *sync.WaitGroup) {
	for i := start; i < fin; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("stop searth, found val", idx)
			return
		default:
			if i == secretVal {
				ch <- i
				fmt.Println("secret found gorutine", idx)
				close(ch)
				return
			}
		}
	}
	wg.Done()

}

func main() {
	secretVal := rand.Intn(10000)
	ch := make(chan int)
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go searthSecretVal(ctx, i*2500, (i+1)*2500, secretVal, i, ch, &wg)
	}

	fmt.Println("found secret val", <-ch)

	cancel()

	go func() {
		wg.Wait()
	}()

	fmt.Println("stopped main")

}
