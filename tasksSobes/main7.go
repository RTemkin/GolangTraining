package main

import "fmt"

func inkrem(sign chan struct{}) {
	for i := 0; i < 100; i++ {
		fmt.Printf("i= %v\n, ", i)
	}

	fmt.Println("Горутина inkrem завершена")
	sign <- struct{}{}
}

func main() {

	sign := make(chan struct{})
	go inkrem(sign)


	<-sign
	fmt.Println("Горутина Main завершена")

}
