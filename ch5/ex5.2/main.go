package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "populate: %v\n", err)
		os.Exit(1)
	}
	dict := make(map[string]int)
	dict = populate(dict, doc)
	fmt.Println(dict)
}

func populate(dict map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		dict[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		dict = populate(dict, c)
	}

	return dict
}
