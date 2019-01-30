package main

import (
	"fmt"
	"os"
	"strings"
)

func expand(s string, f func(string) string) string {
	return strings.Replace(s, "$foo", f("foo"), -1)
}

func replacePart(str string) string {
	return "abcd"
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, `usage: ex5.9 "some_string $foo sdffs"`)
		os.Exit(1)
	}

	fmt.Println(expand(os.Args[1], replacePart))
}
