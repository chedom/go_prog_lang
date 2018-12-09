package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	fileNames := os.Args[1:]
	for _, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error while openning file %v\n", err)
		}
		scanner := bufio.NewScanner(file)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			words[scanner.Text()]++
		}
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "scanner error %v\n", err)
			os.Exit(1)
		}
		file.Close()
	}
	for k, v := range words {
		fmt.Printf("word: %s, count %d", k, v)
	}

}
