//Exercis e 4.3: Re writ e reverse to use an array point er ins tead of a slice.
package main

import "fmt"

func main() {
	var a = [4]int{1, 2, 3, 4}
	reverse(&a)
	fmt.Println(a)
}

func reverse(arr *[4]int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}
