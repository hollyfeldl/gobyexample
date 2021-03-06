package main

// https://gobyexample.com/recursion modified with a for loop

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	for i := 1; i <= 19; i++ {
		fmt.Println("Factoral", i, ":", fact(i))	
	}
	
}