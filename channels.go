package main

// https://gobyexample.com/channels -- played with

import (
	"fmt"
)

func main() {
	messages := make (chan string)
	content := make (chan string)
	var news string
	var msg string

	go func() { messages <- "ping" }()
	go func() { content <- "pong" }()

	fmt.Println("msg before:", msg)
	fmt.Println("news before:", news)

	news = <-content
	fmt.Println("msg #1:", msg)
	fmt.Println("news #1:", news)

	msg = <-messages
	fmt.Println("msg #2:", msg)
	fmt.Println("news #2:", news)

}
