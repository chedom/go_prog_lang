package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

type client struct {
	out chan <- string
	name string
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func broadcaster() {
	clients := make(map[client]struct{})

	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli.out <- msg
			}
		case cli := <-entering:
			clients[cli] = struct{}{}
			go func() { messages <- cli.name }()
		case cli := <- leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{out: ch, name: who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}