package ex11_1

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func Charcount(r io.Reader) (map[rune]int, map[int]int, int) {
	counts := make(map[rune]int)
	utflen := make(map[int]int)
	invalid := 0

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}

	return counts, utflen, invalid
}
