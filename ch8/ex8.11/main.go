package ex8_11

import (
	"context"
	"fmt"
	"io"
	"net/http"
)

func fetch(ctx context.Context, method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.WithContext(ctx)
	return http.DefaultClient.Do(req)
}

func mirroredQuery(urls []string) {
	ctx, cancel := context.WithCancel(context.Background())
	result := make(chan *http.Response)

	for _, url := range urls {
		go func(url string) {
			res, err := fetch(ctx, "GET", url, nil)
			if err != nil {
				return
			}

			select {
			case  result <- res:
				break
			case <- ctx.Done():
				return
			}
		}(url)
	}

	res := <-result
	defer res.Body.Close()

	cancel()


	fmt.Println(res.Request.URL)
}
