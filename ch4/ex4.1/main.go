// Exercis e 4.1: Wr ite a function that counts the number of bits that are dif ferent in two SHA256
// hashes. (See PopCount from Sec tion 2.6.2.)
package main

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	c1 := sha256.Sum256(os.Args[1])
	c2 := sha256.Sum256(os.Args[2])
	fmt.Println(popCount(c1, c2))
}

func popCountByShifting(x uint8) int {
	var count int
	var mask uint8 = 1
	for i := 0; i < 8; i++ {
		if x&mask > 0 {
			count++
		}
		x >>= 1
	}
	return count
}

func popCount(s1, s2 [32]uint8) int {
	var res int
	for k := range s1 {
		res += popCountByShifting(^(s1[k] ^ s2[k]))
	}
	return res
}
