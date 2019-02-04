package main

import "fmt"

func add(a, b int) (result int) {
	type noReturn struct {}
	defer func() {
		switch p:=recover(); p {
		case nil:
			// no panic
		case noReturn{}:
			result = a + b
		default:
			panic(p)
		}
	}()
	panic(noReturn{})
}

func main() {
	fmt.Println(add(12, 13))
}
