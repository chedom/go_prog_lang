package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
)

type StrReader struct {
	str []byte
}

func (s *StrReader) Read(p []byte) (n int, err error) {
	n = copy(p, s.str)

	if n < len(p) {
		err = io.EOF
	}

	s.str = s.str[n:]

	return
}

func NewStringReader(str string) *StrReader {
	return &StrReader{str: []byte(str)}
}

func main() {
	doc, _ := html.Parse(NewStringReader("<html><body><h1>hello</h1></body></html>"))
	fmt.Println(doc.FirstChild.LastChild.FirstChild.FirstChild.Data)
}
