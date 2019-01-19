package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

const ComicURL = "https://xkcd.com/%d/info.0.json"
const LastComicURL = "https://xkcd.com/info.0.json"

type Comic struct {
	Transcript string
	Num        int
}

type NumberIndex map[int]*Comic
type WordIndex map[string]map[int]bool

func getComicUrl(num int) string {
	return fmt.Sprintf(ComicURL, num)
}

func getComicsCount() (int, error) {
	res, err := getComic(LastComicURL)
	if err != nil {
		return 0, err
	}
	return res.Num, nil
}

func getComic(url string) (*Comic, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("comic fetch failed: %s", resp.Status)
	}
	var result Comic
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return &result, nil
}

func filterLetter(word string) string {
	var buffer bytes.Buffer
	for i := 0; i < len(word); {
		r, size := utf8.DecodeRuneInString(word[i:])
		i += size
		if unicode.IsLetter(r) {
			buffer.WriteRune(r)
		}
	}
	return buffer.String()
}

func getWords(str string) []string {
	words := strings.Split(str, " ")
	end := 0
	for i := 0; i < len(words); i++ {
		if filtered := filterLetter(words[i]); len(filtered) != 0 {
			words[end] = strings.ToLower(filtered)
			end++
		}
	}
	return words[:end]
}

func buildIndex() (NumberIndex, WordIndex, error) {
	numIndex := make(NumberIndex)
	wordIndex := make(WordIndex)
	count, err := getComicsCount()
	if err != nil {
		return nil, nil, err
	}
	for i := 1; i <= count; i++ {
		comic, err := getComic(getComicUrl(i))
		if err != nil {
			continue
		}
		numIndex[comic.Num] = comic
		words := getWords(comic.Transcript)
		for _, word := range words {
			if _, ok := wordIndex[word]; !ok {
				wordIndex[word] = make(map[int]bool)
			}
			wordIndex[word][comic.Num] = true
		}
	}
	return numIndex, wordIndex, nil
}

func Build(fileName string) {
	numIndex, wordIndex, err := buildIndex()
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant create file for indexes: %v\n", err)
		return
	}
	defer file.Close()
	enc := gob.NewEncoder(file)
	err = enc.Encode(numIndex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant save num index: %v\n", err)
		return
	}
	err = enc.Encode(wordIndex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant save word index: %v\n", err)
		return
	}
}

func readIndex(fileName string) (NumberIndex, WordIndex, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	var numIndex NumberIndex
	var wordIndex WordIndex
	dec := gob.NewDecoder(file)
	err = dec.Decode(&numIndex)
	if err != nil {
		return nil, nil, err
	}
	err = dec.Decode(&wordIndex)
	if err != nil {
		return nil, nil, err
	}
	return numIndex, wordIndex, nil
}

func searchByQuery(query []string, numIndex NumberIndex, wordIndex WordIndex) []*Comic {
	found := make(map[int]int)
	for _, q := range query {
		for k := range wordIndex[q] {
			found[k]++
		}
	}
	var comics []*Comic
	for k, v := range found {
		if v == len(query) {
			comics = append(comics, numIndex[k])
		}
	}
	return comics
}

func Search(filename string, query []string) {
	numIndex, wordIndex, err := readIndex(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant open index file: %v\n", err)
		return
	}
	comics := searchByQuery(query, numIndex, wordIndex)
	for _, comic := range comics {
		fmt.Printf("Url: %s\n", getComicUrl(comic.Num))
		fmt.Printf("Transcript: %s\n", comic.Transcript)
	}
}

var operation = flag.String("operation", "build", "build|search")
var filename = flag.String("filename", "", "filename for index")

func main() {
	flag.Parse()
	switch *operation {
	case "build":
		Build(*filename)
	case "search":
		Search(*filename, os.Args[3:])
	}
}
