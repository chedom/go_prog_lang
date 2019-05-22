package main

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please input a zap/tar file")
		os.Exit(1)
	}

	files := os.Args[1:]

	for _, filename := range files {
		if strings.HasSuffix(filename, ".zip") {
			handleZip(filename)
		} else if strings.HasPrefix(filename, ".tar") {
			handleTar(filename)
		} else {
			fmt.Fprintf(os.Stderr, "Cant recognize file %s\n", filename)
		}
	}
}

func handleZip(filename string) {
	r, err := zip.OpenReader(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("File of %s:\n", f.Name)
	}
}

func handleTar(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	tr := tar.NewReader(f)

	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // end of archive
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("File of %s:\n", hdr.Name)
	}
}
