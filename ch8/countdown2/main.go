package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	abort := make(chan struct{})

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
			continue
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
}
