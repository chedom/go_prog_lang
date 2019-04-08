package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	if t == nil {
		return ""
	}
	l := t.left.String()
	v := strconv.Itoa(t.value)
	r := t.right.String()
	return fmt.Sprintf("%s %s %s", l, v, r)
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func fill(s []int) *tree {
	t := new(tree)
	for _, v := range s {
		t = add(t, v)
	}
	return t
}

func main() {
	data := make([]int, 50)
	for i := range data {
		data[i] = rand.Int() % 50
	}

	t := fill(data)
	fmt.Println(t)
}
