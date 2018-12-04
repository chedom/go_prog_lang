//Exercis e 4.4: Wr ite a version of rotate that operates in a single pass.
package main

import "fmt"

func main() {
	test := []int{0, 1, 2, 3, 4, 5}
	rotateLeftByN(test, 2)
	fmt.Println(test)
}

func rotateLeftByN(s []int, n int) {
	z := make([]int, 0, len(s))
	for k := range s {
		i := k + n
		if i >= len(s) {
			i -= len(s)
		}
		z = append(z, s[i])
	}
	copy(s, z)
}
