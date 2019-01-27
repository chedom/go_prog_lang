package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var attrs = []string{"href", "src"}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func existInArr(sl []string, s string) bool {
	for _, v := range sl {
		if v == s {
			return true
		}
		continue
	}

	return false
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if existInArr(attrs, a.Key) {
				links = append(links, a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}

	return links
}
