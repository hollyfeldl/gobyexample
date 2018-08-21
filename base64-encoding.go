// https://gobyexample.com/base64-encoding

package main

import(
	b64 "encoding/base64"
	"fmt"
)

func main() {

	data := "dumb://abs.123.com/app&input=var~tmp"
	fmt.Println("Input: ", data)
	fmt.Println()

	sEnc := b64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(sEnc)

	sDec, _ := b64.StdEncoding.DecodeString(sEnc)
	fmt.Println("Std Output: ", string(sDec))
	fmt.Println()

	uEnc := b64.URLEncoding.EncodeToString([]byte(data))
	fmt.Println(uEnc)
	uDec, _ := b64.URLEncoding.DecodeString(uEnc)
	fmt.Println("URL Output: ", string(uDec))

}