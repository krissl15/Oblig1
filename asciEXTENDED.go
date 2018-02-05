package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

//////////////////LØSNING/////////////
func main() {
	src := []byte("20e282ac20c3b720c2be20646f6c6c617220")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n])

}
/////////////////////////LØSNING SLUTT////////////////////
func main() {
	const s = "e282ac"//<-- skriv in bytes i hex her
	decoded, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", decoded)//<-- blir printet ut som tekst

}


package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	src := []byte("Hello Gopher!")//<-- skriv in tekst her

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)//<-- tilsvarende bytes i hex blir printet ut

}
