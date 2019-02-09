package main

import (
	"fmt"
	"io"
	"os"
)

type countWriter struct {
	counter int64
	wrapped io.Writer
}

func (c *countWriter) Write(p []byte) (int, error) {
	n, err := c.wrapped.Write(p)
	c.counter += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := countWriter{wrapped: w}
	return &newWriter, &newWriter.counter
}

func main() {
	newWriter, count := CountingWriter(os.Stdout)
	fmt.Fprintf(newWriter, "test a\n")
	fmt.Println(*count)
	fmt.Fprintf(newWriter, "test b\n")
	fmt.Println(*count)

}
