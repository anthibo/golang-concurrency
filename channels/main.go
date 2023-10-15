package main

import (
	"fmt"
	"strings"
)

func shout(ping, pong chan string) {
	for {
		// receive data from channel
		s, ok := <-ping
		// if closed, the ok is false
		if !ok {
			// do something here
		}

		// send data to a channel
		pong <- fmt.Sprintf("%s!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	fmt.Println("Type something and press Enter, (press q to quit)")

	for {
		// print a prompt
		fmt.Print("-> ")
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		// wait for a response

		response := <-pong

		fmt.Println("Response: ", response)
	}
	fmt.Println("All done, closing channels")
	close(ping)
	close(pong)
}
