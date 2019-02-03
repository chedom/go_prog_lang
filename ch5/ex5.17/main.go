package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var elements []*html.Node
	f := func (n *html.Node) {
		if n.Type == html.ElementNode && contains(name, n.Data) {
			elements = append(elements, n)
		}
	}
	traverse(doc, f)
	return elements
}

func contains(s []string, e string) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func traverse(doc *html.Node, f func(doc *html.Node)) {
	f(doc)

	for c:= doc.FirstChild; c != nil; c = c.NextSibling {
		traverse(c, f)
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cant parse html: %v", err)
		os.Exit(1)
	}
	finds := ElementsByTagName(doc, os.Args[1:]...)
	fmt.Println(finds)
}
