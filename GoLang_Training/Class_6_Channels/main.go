package main

import (
	"fmt"
)

func produce(ch chan int) {
	for i := 1; i <= 3; i++ {
		fmt.Println("Sending:", i)
		ch <- i
	}
	close(ch)
}

func consume(ch chan int) {
	for val := range ch {
		fmt.Println("Received:", val)
	}
}

func main() {
	ch := make(chan int, 2)

	go produce(ch)
	consume(ch)
}

////

/*func send(ch chan string) {
	ch <- "Hello from sender"
}

func main() {
	ch := make(chan string)
	go send(ch)

	msg := <-ch
	fmt.Println(msg)
}*/
