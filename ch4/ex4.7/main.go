//Exercis e 4.7: Mo dif y reverse to reverse the charac ters of a []byte slice that represents a
//UTF-8-enco ded str ing , in place. Can you do it wit hout allocat ing new memor y?
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	a := []byte("abb")
	fmt.Println(string(reverseRune(a)))

}

func reverseRune(s []byte) []byte {
	if len(s) == 0 {
		return s
	}
	_, size := utf8.DecodeRune(s)
	return append(reverseRune(s[size:]), s[:size]...)
}
