package main

import (
	"fmt"
	"time"
)

func listToChan(ch chan int) {
	for {
		// print got a data message
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}
func main() {
	ch := make(chan int, 10)

	go listToChan(ch)

	for i := 0; i <= 100; i++ {
		fmt.Println("Sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel")
	}

	fmt.Println("Done!")
	close(ch)
}
