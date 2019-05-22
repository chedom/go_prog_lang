// ex9.5 tests of performance of ping-ponging goroutines.
package main

import (
	"fmt"
	"time"
)

func main() {
	q := make(chan int)
	var i int64
	start := time.Now()
	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	go func() {
		q <- 1
		for {
			select {
			case <-done:
				return
			case n := <-q:
				i++
				q <- n
			}
		}
	}()
	go func() {
		for {
			select {
			case <-done:
				return
			case n := <-q:
				q <- n
			}
		}
	}()

	<-done
	fmt.Println(float64(i)/float64(time.Since(start))*1e9, "round trips per second")

}
