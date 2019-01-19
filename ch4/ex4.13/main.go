package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
)

const URL = "http://img.omdbapi.com/"

type Movie struct {
	Title  string
	Year   string
	Poster string
}

func generateFilename(movie *Movie) string {
	ext := path.Ext(movie.Poster)
	return fmt.Sprintf("%s_%s%s", movie.Title, movie.Year, ext)
}

func getMovie(query string) (*Movie, error) {
	q := url.QueryEscape(query)
	resp, err := http.Get(URL + "?t=" + q)
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, err
	}
	return &movie, nil
}

func savePoster(movie *Movie) {
	resp, err := http.Get(movie.Poster)
	defer resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant get poster: %v\n", err)
		return
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "getting poster failed: %s", resp.Status)
		return
	}
	filename := generateFilename(movie)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant create file: %s", err)
		return
	}
	defer file.Close()
	if _, err = io.Copy(file, resp.Body); err != nil {
		fmt.Fprintf(os.Stderr, "Cant save to file: %s", err)
		return
	}
}

func main() {
	movieTitle := os.Args[1]
	movie, err := getMovie(movieTitle)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: v\n", err)
		os.Exit(1)
	}
	savePoster(movie)
}
