package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var port = flag.Int("port", 8080, "port for server")

func init() {
	flag.Parse()
}

func main() {
	addr := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	cmd := NewCmd(c)
	cmd.Listen()
}
