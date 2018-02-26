package main

import (
	"bufio"
	"fmt"
	"os"
)

//Foreløpig problem:
//tekst kan ikke gå inn i ch1 fordi det er en byte. Channels kan bare ta i mot typen int.
//Scanner kan bare (såvidt jeg vet) bruke string eller byte. Må fikses eller finne alternativ til scanner.

func main () {

	//ch1 := make(chan int, 2)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		tekst := scanner.Bytes()
		//ch1 <- tekst
		fmt.Println(tekst)
	}

	if scanner.Err() != nil {
		// handle error.
	}
}