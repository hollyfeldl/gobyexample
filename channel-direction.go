package main

// https://gobyexample.com/channel-directions
// modified to tag things and run it though things again

import(
	"fmt"
)

func ping(pings chan<- string, msg string) {

	// tag the ping
	msg += " tag-ping"

	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings

	// tag the pong
	msg = "tag-pong " + msg

	pongs <- msg
}

func main() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed-message")
	pong(pings, pongs)

	// run it though again
	ping(pings, <-pongs)
	pong(pings, pongs)

	fmt.Println(<-pongs)

}

