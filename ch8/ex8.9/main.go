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
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64, sema chan struct{}) {
	defer n.Done()

	for _, entry := range dirents(dir, sema) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes, sema)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string, sema chan struct{}) []os.FileInfo {
	sema <- struct{}{}
	defer func() { <-sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du3: %v\n", err)
		return nil
	}
	return entries
}

func countRoot(root string, n *sync.WaitGroup) {
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	// Traverse the file tree
	fileSizes := make(chan int64)
	var wg sync.WaitGroup

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	go func(root string) {
		defer n.Done()

		var nfiles, nbytes int64

	loop:
		for {
			select {
			case size, ok := <-fileSizes:
				if !ok {
					break loop
				}
				nfiles++
				nbytes += size
			case <-tick:
				printDiskUsage(root, nfiles, nbytes)
			}
		}

		printDiskUsage(root, nfiles, nbytes)

	}(root)

	wg.Add(1)

	// sema is a counting semaphore for limiting concurrncy in dirents
	sema := make(chan struct{}, 10)

	go walkDir(root, &wg, fileSizes, sema)
}

func main() {
	// Determine the initial directories.
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var n sync.WaitGroup
	n.Add(len(roots))

	for _, root := range roots {
		// Print the result periodically.
		go countRoot(root, &n)
	}

	n.Wait()
}

func printDiskUsage(dir string, nfiles, nbytes int64) {
	fmt.Printf("%s - %d files %.1f MB\n", dir, nfiles, float64(nbytes)/1e6)
}
