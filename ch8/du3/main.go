package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")

// walkDir recursively walk the file tree rooted at dir
// and sends the size of each found file on fileSize
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan <- int64) {
	defer n.Done()

	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// sema is a counting semaphore for limiting concurrncy in dirents
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// Traverse the file tree
	fileSizes := make(chan int64)

	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, &wg, fileSizes)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print the result periodically.
	var tick <- chan time.Time
	if *verbose {
		tick = time.Tick(500 *time.Millisecond)
	}

	var nfiles, nbytes int64

	loop:
		for {
			select {
			case size, ok := <- fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <- tick:
				printDiskUsage(nfiles, nbytes)
			}
		}


	printDiskUsage(nfiles, nbytes)
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}
