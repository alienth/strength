package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	foo, _ := rand.Int(rand.Reader, big.NewInt(int64(len(wordlist))))
	fmt.Println(wordlist[foo.Int64()])
}
