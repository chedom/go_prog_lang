package xkcd

const ComicURL = "https://xkcd.com/%d/info.0.json"
const LastComicURL = "https://xkcd.com/info.0.json"

type Comic struct {
	Transcript string
	Num        int
}

type NumberIndex map[int]Comic
type WordIndex map[string]map[string]bool
