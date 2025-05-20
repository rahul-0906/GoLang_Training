package main

import (
	"fmt"
)

func printOdd(oddCh, evenCh chan int, done chan bool) {
	for i := 1; i <= 9; i += 2 {
		<-oddCh // wait for signal to print odd
		fmt.Println("Odd:", i)
		evenCh <- 1 // signal even printer
	}
	done <- true
}

func printEven(oddCh, evenCh chan int, done chan bool) {
	for i := 2; i <= 10; i += 2 {
		<-evenCh // wait for signal to print even
		fmt.Println("Even:", i)
		oddCh <- 1 // signal odd printer
	}
	done <- true
}

func main() {
	oddCh := make(chan int)  // unbuffered
	evenCh := make(chan int) // unbuffered
	done := make(chan bool)  // to signal completion

	go printOdd(oddCh, evenCh, done)
	go printEven(oddCh, evenCh, done)

	oddCh <- 1 // start with odd printer

	// Wait for both goroutines to finish
	<-done
	<-done
}
