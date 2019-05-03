package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var addr = flag.String("addr", "localhost:8080", "server address")

func main() {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("mustCopy: %v\n", err)
	}
}
