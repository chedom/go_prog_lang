package main

import (
	"os"

	"github.com/chedom/go_prog_lang/ch5/links"
)

func crawl(url string) []string {
	if list, err := links.Extract(url); err != nil {
		return nil
	} else {
		return list
	}
}

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)
	unsavedLinks := make(chan string)

	var n int

	n++
	go func() { worklist <- os.Args[1:] }()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unsavedLinks {
				save(link)
			}
		}()
	}

	seen := make(map[string]bool)

	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				unseenLinks <- link
				unsavedLinks <- link
			}
		}
	}
}
