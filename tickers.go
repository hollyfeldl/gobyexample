// https://gobyexample.com/tickers

package main

import(
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 250)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	} ()

	time.Sleep(time.Millisecond * 1600)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
