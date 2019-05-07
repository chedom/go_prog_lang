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
	ctx, cancelFunc := context.WithCancel(context.Background())
	out := make(chan string)
	in := make(chan string)

	go clientWriter(conn, out)
	go clientReader(conn, in)

	out <- "Write you name"

	go closer(ctx, conn, in)
	var who string
	who = <- in

	cli := client{out: out, name: who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli


	for m := range in {
		messages <- who + ": " + m
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
	cancelFunc()
}

func clientWriter(conn net.Conn, ch <- chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientReader(conn net.Conn, ch chan <- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}

func closer(ctx context.Context, conn net.Conn, sendEvents <- chan string) {
	go func() {
		for {
			select {
			case <- ctx.Done():
				return
			case <- time.After(5 * time.Minute):
				conn.Close()
				return
			case <-sendEvents:
				continue
			}
		}
	}()
}