package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

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

func buildIndex(fileName string) {
	numberIndex := make(NumberIndex)
	wordIndex := make(WordIndex)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant create file for indexes: %v\n", err)
		return
	}
	defer file.Close()
	stringifyNum, err := json.Marshal(numberIndex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant stringify number index: %v\n", err)
		return
	}
	if _, err := file.Write(stringifyNum); err != nil {
		fmt.Fprintf(os.Stderr, "Cant save num index: %v\n", err)
		return
	}
	stringifyWord, err := json.Marshal(wordIndex)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cant stringify word index: %v\n", err)
		return
	}
	if _, err := file.Write(stringifyWord); err != nil {
		fmt.Fprintf(os.Stderr, "Cant save word index: %v\n", err)
		return
	}

}
