package main

func SplitChannel[T any](inputCh <-chan T, n int) []<-chan T {
	outputChs := make([]chan T, n)
	for i := 0; i < n; i++ {
		outputChs[i] = make(chan T)
	}

	idx := 0
	for value := range inputCh {
		outputChs[idx] <- value
		idx = (idx + 1) % n
	}

	for _, ch := range outputChs {
		close(ch)
	}

}

func main() {

}
