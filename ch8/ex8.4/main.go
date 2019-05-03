package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var port = flag.Int("port", 8080, " port for tcp server")

func main() {
	flag.Parse()
	addr := fmt.Sprintf("localhost:%d", *port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Accept error: %v\n", err)
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	var wg sync.WaitGroup

	for input.Scan() {
		wg.Add(1)

		go func(s string) {
			defer wg.Done()

			echo(c, s, 1*time.Second)
		}(input.Text())
	}

	wg.Wait()
	if conn, ok := c.(*net.TCPConn); ok {
		conn.CloseWrite()
		return
	}
	c.Close()
}
