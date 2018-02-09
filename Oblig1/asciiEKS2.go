package main

import (
	"encoding/hex"
	"fmt"
	"log"
)
/**
* byteToHex vil gjøre om symbolene en skriver i "src" til hex-verdier. 
*/


func byteToHex() {
	src := []byte(" € ÷ ¾ dollar ")

	dst := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(dst, src)

	fmt.Printf("%s\n", dst)
}


/**
*main funksjonen vil gjøre om hex-verdien fra forrige funksjon til symboler (Slik oppgaven ber om).
* Å bruke disse to funksjonene om hverandre gir en grei løsning på oppgaven, men krever ingen forståelse av unicode. 
*/
func main() {
	src := []byte("20e282ac20c3b720c2be20646f6c6c617220") // Bruker en annen f

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", dst[:n]) // output:  € ÷ ¾ dollar 

}
