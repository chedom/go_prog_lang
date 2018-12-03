// Exercis e 3.10: Wr ite a non-rec ursive version of comma , using bytes.Buffer instead of str ing
// conc atenation.
package main

import (
	"bytes"
)

func main() {
	comma("12345")
}

func comma(str string) string {
	var buf bytes.Buffer
	n := len(str)
	start := n % 3
	if start == 0 {
		start = 3
	}
	buf.WriteString(str[:start])
	for i := start; start < n; i += 3 {
		buf.WriteRune(',')
		buf.WriteString(str[start : start+3])
	}

	return buf.String()
}
