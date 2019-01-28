package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int

func nodeHasChildren(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode || c.Type == html.TextNode {
			return true
		}
		continue
	}
	return false
}

func startElement(n *html.Node) {
	if n.Type == html.CommentNode {
		fmt.Printf("%*s//%s\n", depth*2, "", n.Data)
		return
	}
	if n.Type == html.TextNode {
		fmt.Printf("%*s %s\n", depth*2, "", n.Data)
		return
	}
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s ", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Printf("%s=%s ", a.Key, a.Val)
		}
		if nodeHasChildren(n) {
			fmt.Printf(">\n")
		} else {
			fmt.Printf("/> \n")
		}
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if nodeHasChildren(n) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline2: parse: %v", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)
}
