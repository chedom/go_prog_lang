// Exercise 1.4: Mo dif y dup2 to print the names of all files in which each dup lic ated line occ urs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	lineInFiles := make(map[string][]string)
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, lineInFiles)
		f.Close()
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, lineInFiles[line])
		}
	}
}

func strInSlice(str string, s []string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, lineInFiles map[string][]string) {
	fileName := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !strInSlice(line, lineInFiles[line]) {
			lineInFiles[line] = append(lineInFiles[line], fileName)
		}
	}
}
