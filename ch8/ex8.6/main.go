package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/chedom/go_prog_lang/ch5/links"
)

var depth = flag.Int("depth", 3, "depth-limiting to the concurrent")

var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)

	tokens <- struct{}{} // acquire a token
	list, err := links.Extract(url)
	<-tokens //release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	flag.Parse()

	type item struct {
		list  []string
		depth int
	}
	worklist := make(chan item)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments
	n++
	go func() {
		i := item{list: flag.Args(), depth: 3}
		worklist <- i
	}()

	// Crawl the web concurrently
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		i := <-worklist

		if i.depth == 0 {
			continue
		}
		fmt.Println(">>>>>>", i.depth)
		for _, link := range i.list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(i item) {
					i = item{depth: i.depth - 1, list: crawl(link)}
					worklist <- i
				}(i)

			}
		}
	}

}
