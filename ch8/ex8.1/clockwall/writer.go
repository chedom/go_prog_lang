package main

import (
	"log"
	"os"
)

type withCity struct {
	city string
}

func (c *withCity) Write(b []byte) (int, error) {
	city := []byte(c.city + ": ")
	res := append(city, b...)
	if n, err := os.Stdout.Write(res); err == nil {
		return n - len(city), err
	} else {
		log.Printf("write to io: %v\n", err)
		return n, err
	}

}
