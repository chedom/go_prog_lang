package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

type client struct {
	out  chan<- string
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
		case cli := <-leaving:
			delete(clients, cli)
			close(cli.out)
		}
	}
}

func handleConn(conn net.Conn) {
	ctx, cancelFunc := context.WithCancel(context.Background())
	ch := make(chan string)
	sendEvents := make(chan struct{})

	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	cli := client{out: ch, name: who}
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	input := bufio.NewScanner(conn)
	go closer(ctx, conn, sendEvents)

	for input.Scan() {
		sendEvents <- struct{}{}
		messages <- who + ": " + input.Text()
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
	cancelFunc()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func closer(ctx context.Context, conn net.Conn, sendEvents <-chan struct{}) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(5 * time.Minute):
				conn.Close()
				return
			case <-sendEvents:
				continue
			}
		}
	}()
}
