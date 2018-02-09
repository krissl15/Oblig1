package main

import (
	"testing"
)

func asciiTest(j string) bool {
	for _, i := range j {
		if i < 127 || i > 255  {
			return true
		}
	}
	return false
}

func TestExtendedASCII(t *testing.T) {

	asciiR := []rune{34, 8364, 32, 190, 32, 247, 32, 100, 111, 108, 108, 97, 114, 34}
	asciiS := string(asciiR)
	asciiByteSlice := []byte(asciiS)

	g := ExtendedASCIIText(asciiByteSlice)


	for _, d := range g { //<-- Skriv in strings som skal testes her.
		h := string(d)
		if asciiTest(h) == true {

			t.Errorf("%s %v", h, "This string is  NOT extended ASCII")
		} else {
			t.Errorf("%s %v", h, "This string IS from the extended ASCII")

		}
	}

}
