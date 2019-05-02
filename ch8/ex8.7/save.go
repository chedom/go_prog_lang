package main

import (
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"golang.org/x/net/html"
)

var originalHost string

func save(rawurl string) error {
	parsedURL, err := url.Parse(rawurl)

	if err != nil {
		return err
	}

	if originalHost == "" {
		originalHost = parsedURL.Host
	}

	if originalHost != parsedURL.Host {
		return nil
	}

	resp, err := http.Get(rawurl)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	dirname := filepath.Join(".", "test", parsedURL.Host)
	var filename string
	if filepath.Ext(parsedURL.Path) == "" {
		dirname = filepath.Join(dirname, parsedURL.Path)
		filename = filepath.Join(dirname, "index.html")
	} else {
		dirname = filepath.Join(dirname, filepath.Dir(parsedURL.Path))
		filename = filepath.Join(dirname, filepath.Base(parsedURL.Path))
	}

	if err := os.MkdirAll(dirname, os.ModePerm); err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}


	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil
	}

	preprocess(doc, originalHost)

	if err = html.Render(file, doc); err != nil {
		return err
	}

	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}


