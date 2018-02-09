package main

import "fmt"

func main() {
	t1 := []byte{34, 190, 32, 247, 32, 100, 111, 108, 108, 97, 114, 34}
	for i := 0; i < len(t1); i++ {
		fmt.Printf("%c", t1[i])
	}

}
