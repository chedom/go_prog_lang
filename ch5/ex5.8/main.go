package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if res := forEachNode(c, pre, post); res != nil {
			return res
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}

	return nil
}

func ElementById(node *html.Node, id string) *html.Node {
	f := func(n *html.Node) bool {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if a.Key == "id" {
					return !(a.Val == id)
				}
			}
		}
		return true
	}
	return forEachNode(node, f, nil)
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: parse: %v", err)
		os.Exit(1)
	}
	id := os.Args[1]
	res := ElementById(doc, id)
	fmt.Println(res)
}
