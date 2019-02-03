package main

import "fmt"

func max(max int, vals ...int) int {
	for _, v := range vals {
		if v > max {
			max = v
		}
	}

	return max
}

func min(min int, vals ...int) int {
	for _, v := range vals {
		if v < min {
			min = v
		}
	}

	return min
}

func main() {
	fmt.Println(max(3,35,78,23,2,45,77,8))
	fmt.Println(min(3,35,78,23,2,45,77,8))
}
