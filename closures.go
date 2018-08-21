package main

// https://gobyexample.com/closures

import (
	"fmt"
)

func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	nextInt := intSeq()

	fmt.Println("nextInt #1: ", nextInt())
	fmt.Println("nextInt #2: ", nextInt())
	fmt.Println("nextInt #3: ", nextInt())

	newInts := intSeq()
	fmt.Println("newInts #1: ", newInts())

}