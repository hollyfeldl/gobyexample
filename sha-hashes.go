// https://gobyexample.com/sha1-hashes

package main

import(
	"crypto/sha1"
	"crypto/sha256"
	"golang.org/x/crypto/sha3"
	"fmt"
)

func main() {

	s := "The quick brown fox jumps over the lazy dog"
	h := sha1.New()
	h_sha2 := sha256.New()
	h_sha3 := sha3.New256()

	h.Write([]byte(s))
	bs := h.Sum(nil)

	h_sha2.Write([]byte(s))
	bs_sha2 := h_sha2.Sum(nil) // Sum here is the method call of the Hash type and not the Sum256 of the sha256 pkg

	h_sha3.Write([]byte(s))
	bs_sha3 := h_sha3.Sum(nil) // ditto

	fmt.Println(s)
	fmt.Printf("SHA1: %x\n", bs)
	fmt.Printf("SHA2-256: %x\n", bs_sha2)
	fmt.Printf("SHA3-256: %x\n", bs_sha3)

}