package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

func connect(city, addr string, wg sync.WaitGroup) error {
	defer wg.Done()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("error connect: %v\n", err)
		return err
	}
	return copy(&withCity{city:city}, conn)

}


func copy(dst io.Writer, src io.Reader) error {

	if _, err := io.Copy(dst, src); err != nil {
		log.Printf("error copy: %v\n", err)
		return err
	}

	return nil
}

func main() {
	args := os.Args[1:]
	var wg sync.WaitGroup
	wg.Add(len(args))
	for _, v := range args {
		a := strings.Split(v, "=")
		go connect(a[0], a[1], wg)
	}
	wg.Wait()
}
