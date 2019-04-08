package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
)

var attrReg = regexp.MustCompile(`^(\w*)\[(\w+)=(\w+)\]$`)

func main() {
	dec := xml.NewDecoder(os.Stdin)
	queries := parseArgs(os.Args[1:])

	var stack []xml.StartElement
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok) // push
		case xml.EndElement:
			stack = stack[:len(stack) - 1] // pop
		case xml.CharData:
			if containsAll(stack, queries) {
				fmt.Printf("%s: %s\n", formatStack(stack), tok)
			}

		}
	}
}

func formatStack(stack []xml.StartElement) string {
	names := make([]string, 0, len(stack))
	for _, v := range stack {
		names = append(names, v.Name.Local)
	}

	return strings.Join(names, " ")
}

type Query struct {
	el string
	attr *Pair
}

func parseArgs(strs []string) []Query {
	qs := make([]Query, 0, len(strs))
	for _, v := range strs {
		qs = append(qs, decodeQuery(v))
	}

	return qs
}

func containsAll(x []xml.StartElement, y []Query) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}

		if containToken(x[0], y[0]) {
			y = y[1:]
		}
		x = x[1:]
	}

	return false
}

func containToken(token xml.StartElement, q Query) bool {
	if q.el != "" {
		if token.Name.Local != q.el {
			return false
		}
	}
	if q.attr != nil {
		for _, a := range token.Attr {
			if a.Name.Local == q.attr.name {
				if a.Value != q.attr.val {
					return  false
				}
			}
		}
	}

	return true
}

type Pair struct {
	name string
	val string
}

func decodeQuery(q string) Query {
	match := attrReg.FindAllSubmatch([]byte(q), -1)
	if  match == nil {
		return Query{ el:q}
	}

	sRes := match[0]

	if len(sRes) != 4 {
		fmt.Fprintf(os.Stderr, "decodeQuery: %q\n", match)
		return Query{ el:q}
	}

	el := string(sRes[1])

	attr := &Pair{
		name: string(sRes[2]),
		val: string(sRes[3]),
	}

	return Query{ el:el, attr: attr}
}