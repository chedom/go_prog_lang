package main

import (
	"fmt"
	"log"
	"os"
)

func crawl(cancel chan struct{},url string) []string {
	fmt.Println(url)
	list, err := Extract(cancel, url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	cancel := make(chan struct{})

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(cancel)
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for {
				select {
				case link, ok := <- unseenLinks:
					if !ok {
						close(worklist)
						return
					}
					foundLinks := crawl(cancel, link)
					go func() {
						select {
						case worklist <- foundLinks:
							break
						case <- cancel:
							return
						}
					}()
				case <- cancel:
					for range unseenLinks {
						// do nothing
					}
					close(worklist)
				}
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				select {
				case unseenLinks <- link:
					break
				case <- cancel:
					for range worklist {

					}
					close(unseenLinks)
				}
			}
		}
	}

}
