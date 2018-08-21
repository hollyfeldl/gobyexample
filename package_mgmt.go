// a go package example 
// extends go by example

package main

import (
	"fmt"
	"github.com/hollyfeldl/gobyexample/plus"
)

func main() {

	res := plus.Plus(4, 5)
	fmt.Println("4+5 =", res)

	res = plus.PlusPlus(4, 5, 6)
	fmt.Println("4+5+6 =", res)

}
