// Exercis e 1.2: Mo dif y the echo prog ram to print the index and value of each of its arguments,
// on e per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	for k, v := range os.Args[1:] {
		fmt.Println(k, v)
	}
}
