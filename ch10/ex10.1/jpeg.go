package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"
)

var outFormat = flag.String("f", "jpeg", "file format")

func main() {
	flag.Parse()
	if err := toFormat(os.Stdin, os.Stdout, *outFormat); err != nil {
		fmt.Fprintf(os.Stderr, "jpeg: %v\n", err)
		os.Exit(1)
	}
}

func toFormat(in io.Reader, out io.Writer, format string) error {
	format = strings.ToLower(format)
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, nil)

	}

	return errors.New("cant recognize format")
}
