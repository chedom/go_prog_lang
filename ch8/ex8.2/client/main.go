package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var addr = flag.String("addr", "localhost:8080", "address of server")

func main() {
	conn, err := net.Dial("tcp", *addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	go copy(os.Stdout, conn)

	sendCmd(conn)

}

func copy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("copy data: %v\n", err)
	}
}

func sendCmd(conn net.Conn) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		line := scanner.Text()
		words := strings.Split(line, " ")
		if len(words) == 0 {
			continue
		}
		fmt.Fprintf(conn, line+"\n")
		if words[0] == "close" {
			break
		}
	}
}
