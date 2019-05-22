package main

import (
	"bufio"
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
	out := make(chan string)
	in := make(chan string)

	go clientWriter(conn, out)
	go clientReader(conn, in)

	out <- "Write you name"

	var who string
	select {
	case name := <-in:
		who = name
	case <-time.After(5 * time.Minute):
		conn.Close()
		return
	}

	cli := client{out: out, name: who}
	out <- "You are " + who
	messages <- who + " has arrived"
	entering <- cli

	for m := range in {
		messages <- who + ": " + m
	}

Loop:
	for {
		select {
		case msg := <-in:
			messages <- who + ": " + msg
			continue
		case <-time.After(5 * time.Minute):
			break Loop
		}
	}

	leaving <- cli
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

func clientReader(conn net.Conn, ch chan<- string) {
	input := bufio.NewScanner(conn)
	for input.Scan() {
		ch <- input.Text()
	}
}
