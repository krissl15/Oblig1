package main

import "testing"

func asciiTest(j string) bool {
	for _, i := range j {
		if i < 127 || i> 255 {
			return false
		}
	}
	return true
}

func testExtendedASCII(t * testing.B){
	g := ExtendedASCIIText()
	if asciiTest(g) == false {
		t.Errorf("fail")
	}
}





