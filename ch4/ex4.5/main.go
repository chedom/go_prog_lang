//Exercis e 4.5: Wr ite an in-place function to eliminate adj acent dup lic ates in a []string slice.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicateAdjacent([]string{"d", "w", "w", "d"}))
}

func removeDuplicateAdjacent(s []string) []string {
	i := 0
	for _, v := range s {
		if s[i] != v {
			i++
			s[i] = v
		}
	}
	return s[:i+1]
}
