package main

import (
	"fmt"
	"os"

	"github.com/chedom/go_prog_lang/ch4/ex4.11/issuesutil"
)

func main() {
	args := os.Args[1:]
	res, err := issuesutil.GetIssue(args[0], args[1], args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	fmt.Printf("result: %v\n", res)
}
