package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

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
