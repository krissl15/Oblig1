package main

import (
	"testing"
	"strings"
)

func asciiTest(j string) bool {
	for _, i := range j {
		if i < 127 || i > 255 {
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

	testArray := strings.Fields(g)
	for _, d := range testArray {
		if asciiTest(g) == true {
			t.Errorf("Some of these strings are not extended ASCII")
		} else {
			t.Errorf("These strings are from the extended ASCII")

		}
	}

}
