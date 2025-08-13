package main

import (
	"fmt"
	"time"
)

func player1(table chan int) {
	for {
		ball := <- table
		ball++
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Ping")
		table <-ball

	}
}

func player2(table chan int) {
	for {
		ball := <- table
		ball++
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Pong")
		table <-ball

	}

}

func main() {

	table := make(chan int)
	Ball := 0

	go player2(table)
	go player1(table)

	table <- Ball
	time.Sleep(1 * time.Second)
	<-table

}