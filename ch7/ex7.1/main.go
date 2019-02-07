package main

import (
	"bufio"
	"bytes"
	"fmt"
)

type WordCounter int

func (w *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*w++
	}
	if err:= scanner.Err(); err != nil {
		return 0, err
	}
	return len(p), nil
}

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*l++
	}
	if err:= scanner.Err(); err != nil {
		return 0, nil
	}
	return len(p), nil
}

func main() {
	var w WordCounter
	var l LineCounter
	t := `there
		are
		5 
		lines
		.`
	w.Write([]byte("There are 4 words"))
	l.Write([]byte(t))
	fmt.Println(w, l)
}
