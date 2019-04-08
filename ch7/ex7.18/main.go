package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Node interface{}

type CharData string
type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}

func (e *Element) String() string {
	b := &bytes.Buffer{}
	visit(e, b, 0)
	return b.String()
}

func visit(n Node, w io.Writer, depth int) {
	switch n := n.(type) {
	case *Element:
		fmt.Fprintf(w, "%*s%s %s\n", depth*2, "", n.Type.Local, n.Attr)
		for _, c := range n.Children {
			visit(c, w, depth+1)
		}
	case CharData:
		fmt.Fprintf(w, "%*s%q\n", depth*2, "", n)
	default:
		panic(fmt.Sprintf("got %T", n))
	}
}

func parse(r io.Reader) (Node, error) {
	dec := xml.NewDecoder(r)
	var stack []*Element
	var tree Node

	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			return nil, err
		}

		switch tok := tok.(type) {
		case xml.StartElement:
			n := &Element{Type: tok.Name, Attr: tok.Attr, Children: nil}
			if len(stack) == 0 {
				tree = n
			} else {
				last := stack[len(stack)-1]
				last.Children = append(last.Children, n)
			}
			stack = append(stack, n) // push

		case xml.EndElement:
			stack = stack[:len(stack)-1]
		case xml.CharData:
			last := stack[len(stack)-1]
			last.Children = append(last.Children, tok)
		}
	}

	return tree, nil
}

func main() {
	node, err := parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	fmt.Println(node)
}
