package main

import (
	"encoding/hex"
	"fmt"
	"log"
)
/**
func stringToHex() {
	src := []byte(" € ÷ ¾ dollar ")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)
}
*/


func hexToString() {
	src := []byte("20e282ac20c3b720c2be20646f6c6c617220") // Bruker en annen f

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])

}
