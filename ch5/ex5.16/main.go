package main

import "fmt"

func Join(sep string, vals ...string) string {
	var result string
	for _, val := range vals {
		if result != "" {
			result += sep
		}
		result += val
	}
	return result
}

func main () {
	fmt.Println(Join("-", "a", "b", "c", "d", "e"))
}