package main

import "fmt"

func ExtendedASCIIText(j string) bool {
	for _, i := range j {
		if i > 127 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(ExtendedASCIIText("xyz"))
	fmt.Println(ExtendedASCIIText("øæå"))
}
