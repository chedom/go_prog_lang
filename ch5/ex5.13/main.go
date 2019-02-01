package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/chedom/go_prog_lang/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				fmt.Println("1", item)
				seen[item] = true
				save(item)
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var originURL string

func crawl(url string) []string {
	fmt.Println("2", url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func save(rawURL string) error {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	if originURL == "" {
		originURL = parsedURL.Host
	}

	if originURL != parsedURL.Host {
		return nil
	}

	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed fatching %s: status: %s", rawURL, resp.Status)
	}

	dirname := filepath.Join(".", "test", parsedURL.Host)
	var filename string
	if filepath.Ext(parsedURL.Path) == "" {
		dirname = filepath.Join(dirname, parsedURL.Path)
		filename = filepath.Join(dirname, "index.html")
	} else {
		dirname = filepath.Join(dirname, filepath.Dir(parsedURL.Path))
		filename = filepath.Join(dirname, filepath.Base(parsedURL.Path))
	}

	if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	if _, err := io.Copy(file, resp.Body); err != nil {
		return err
	}

	if err := resp.Body.Close(); err != nil {
		return err
	}
	if err := file.Close(); err != nil {
		return err
	}
	return nil
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
