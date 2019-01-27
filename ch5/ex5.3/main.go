package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "content: %v\n", err)
		os.Exit(1)
	}
	content(doc)
}

func content(n *html.Node) {

	if n.Data == "script" || n.Data == "style" {
		return
	}

	if n.Type == html.TextNode {
		fmt.Println(n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		content(c)
	}
}
