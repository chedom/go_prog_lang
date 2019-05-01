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
		mustCopy(os.Stdout, conn)
		log.Println("Done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)

	if conn, ok := conn.(*net.TCPConn); ok {
		conn.CloseWrite()
	} else {
		conn.Close()
	}

	<- done
}


func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("mustCopy err: %v\n", err)
	}
}
