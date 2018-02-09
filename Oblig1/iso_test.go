package main

import (
	"testing"
)

func asciiTest(j string) bool {
	for _, i := range j {
		if i < 127 || i > 255 { //<-- Sjekker om codepointet er mellom 127 og 255, altså extended ASCII
			return true
		}
	}
	return false
}

func TestExtendedASCII(t *testing.T) {

/**	asciiR := []rune{34, 8364, 32, 190, 32, 247, 32, 100, 111, 108, 108, 97, 114, 34}	//<-- Rune pga €.
	asciiS := string(asciiR)	//<-- fra rune til string
	asciiByteSlice := []byte(asciiS)	//<-- fra string til []byte

	g := ExtendedASCIIText(asciiByteSlice) //<-- bruker []byte som parameter 

*/

	for _, d := range "æ" { //<-- velg test-string. Fjern "/**"  over og bruk g for å teste tegn fra 4B.
		h := string(d)
		if asciiTest(h) == true {

			t.Errorf("%s %v", h, "This string is  NOT extended ASCII")
		} else {
			t.Errorf("%s %v", h, "This string IS from the extended ASCII")

		}
	}

}
