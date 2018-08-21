package main

// https://gobyexample.com/channel-buffering

import (
	"fmt"
)

func main() {
	
	messages := make(chan string, 3)

	messages <- "buffered"
	messages <- "channel"
	messages <- "overflow"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)

}