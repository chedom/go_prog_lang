// Exercis e 1.8: Mo dif y fetch to add the prefix http:// to each argument URL if it is missing.
// Yo u mig ht want to use strings.HasPrefix .
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		resp, err := http.Get(formatUrl(url))
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: readinf %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

func formatUrl(url string) string {
	prefix := "http://"
	if strings.HasPrefix(url, prefix) {
		return url
	}
	return prefix + url
}
