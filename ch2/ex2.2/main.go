package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/chedom/go_prog_lang/ch2/ex2.2/conv"
)

func processArg(arg string) {
	v, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conv: %v\n", err)
		os.Exit(1)
	}

	fahr := conv.Fahrenheit(v)
	c := conv.Celsius(v)
	fmt.Printf("%s = %s, %s = %s\n", fahr, conv.FToC(fahr), c, conv.CToF(c))

	feet := conv.Feet(v)
	m := conv.Meter(v)
	fmt.Printf("%s = %s, %s = %s\n", feet, conv.FToM(feet), m, conv.MToF(m))

	p := conv.Pound(v)
	k := conv.Kilogram(v)
	fmt.Printf("%s = %s, %s = %s\n", p, conv.PToK(p), k, conv.KToP(k))
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			processArg(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			processArg(scan.Text())
		}
	}
}
