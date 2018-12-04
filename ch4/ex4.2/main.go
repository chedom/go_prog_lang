//Exercis e 4.2: Wr ite a program that prints the SHA256 hash of its stand ard inp ut by defau lt but
//supp orts a command-line flag to print the SHA384 or SHA512 hash ins tead.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var methodStr string
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&methodStr, "m", "256", "sha method")
	mySet.Parse(os.Args[1:])
	var result []byte
	switch os.Args[1] {
	case "256":
		r := sha512.Sum384([]byte(os.Args[2]))
		result = r[:]
	case "s2":
		r := sha512.Sum512([]byte(os.Args[2]))
		result = r[:]
	default:
		r := sha256.Sum256([]byte(os.Args[2]))
		result = r[:]
	}
	printHash(result)
}

func printHash(hash []byte) {
	for _, v := range hash {
		fmt.Printf("%X", v)
	}
	fmt.Println()
}
