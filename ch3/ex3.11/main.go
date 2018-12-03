// Exercis e 3.11: En hance comma so that it deals cor rec tly wit h floating-p oint numbers and an
// opt ion al sig n.
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, v := range os.Args[1:] {
		fmt.Printf("%s - %s\n", v, comma(v))
	}
}

func comma(str string) string {
	var buf bytes.Buffer
	var mantissaStart, mantissaEnd int
	if str[0] == '+' || str[0] == '-' {
		buf.WriteByte(str[0])
		mantissaStart = 1
	}
	mantissaEnd = strings.Index(str, ".")
	if mantissaEnd < 0 {
		mantissaEnd = len(str)
	}
	newStr := str[mantissaStart:mantissaEnd]
	pre := len(newStr) % 3
	if pre == 0 {
		pre = 3
	}
	buf.WriteString(newStr[:pre])
	for i := pre; i < len(newStr); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(newStr[i : i+3])
	}
	if mantissaEnd != len(str) {
		buf.WriteString(str[mantissaEnd:])
	}
	return buf.String()
}
