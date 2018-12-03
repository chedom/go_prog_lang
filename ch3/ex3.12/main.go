// Exercis e 3.12: Wr ite a function that rep orts whether two str ings are anagrams of each other,
// that is, the y cont ain the same letters in a dif ferent order.
package main

import (
	"bytes"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	isAnagrama(os.Args[1], os.Args[2])
}

func isAnagrama(str1, str2 string) bool {
	str1 = filterLetters(str1)
	str2 = filterLetters(str2)
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); {
		r1, size := utf8.DecodeRuneInString(str1[i:])
		i += size
		if r2, _ := utf8.DecodeRuneInString(str2[len(str2)-i:]); r1 != r2 {
			return false
		}
	}
	return true

}

func filterLetters(str string) string {
	var buf bytes.Buffer
	for i := 0; i < len(str); {
		r, size := utf8.DecodeRuneInString(str[i:])
		i += size
		if !unicode.IsLetter(r) {
			continue
		}
		buf.WriteRune(unicode.ToUpper(r))
	}
	return buf.String()
}
