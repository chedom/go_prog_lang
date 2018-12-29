//Exercis e 4.6: Wr ite an in-place function that squ ashes each run of adj acent Unico de sp aces
//(s ee unicode.IsSpace ) in a UTF-8-enco ded []byte slice into a single ASCII space.
package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	b := []byte("l\nbc\r  \n\rdef")
	fmt.Printf("%q\n", squashSpaces(b))
	fmt.Printf("%q\n", b)
}

func squashSpaces(s []byte) []byte {
	n := 0
	for i := 0; i < len(s); {
		v, size := utf8.DecodeRune(s[i:])
		if !unicode.IsSpace(v) {
			utf8.EncodeRune(s[n:], v)
			n += size
		} else {
			s[n] = ' '
			n++
		}
		i += size
	}
	return s[:n]
}
