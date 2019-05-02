package main

import (
	"golang.org/x/net/html"
	"net/url"
)

func preprocess(n *html.Node, originalHost string) {
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for i, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				l, err := url.Parse(a.Val)
				if err != nil {
					continue
				}

				if l.Host == originalHost {
					l.Host = "mirror." + l.Host
					a.Val = l.String()
					n.Attr[i] = a
				}

			}
		}
	}

	forEachNode(n, visitNode, nil)
}

func forEachNode(node *html.Node, pre, post func(*html.Node)) {
	if pre != nil {
		pre(node)
	}

	for c := node.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(node)
	}
}