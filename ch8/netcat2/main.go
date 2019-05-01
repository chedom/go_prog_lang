package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

var addr = flag.String("addr", "localhost:8080", "address of tcp server")

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("mustCopy: %v\n", err)
	}
}
