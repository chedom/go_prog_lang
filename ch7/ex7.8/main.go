package multisort

import (
	"sort"
	"strings"
	"time"
)

type Track struct {
	Title string
	Artist string
	Album string
	Year int
	Length time.Duration
}

// define sort function
type LessFunc func(i, j *Track) bool
func sortByTitle(i, j *Track) bool { return i.Title < j.Title }
func sortByArtist(i, j *Track) bool { return i.Artist < j.Artist }
func sortByAlbum(i, j *Track) bool { return i.Album < j.Album }
func sortByYear(i, j *Track) bool { return i.Year < j.Year }
func sortByLength(i, j *Track) bool { return i.Length < j.Length }

var sortDict = map[string]LessFunc{
	"title": sortByTitle,
	"artist": sortByArtist,
	"album": sortByAlbum,
	"year": sortByYear,
	"length": sortByLength,
}

type MultiSort struct {
	t []*Track
	compares []LessFunc
}

func (m *MultiSort) Len() int { return len(m.t) }
func (m *MultiSort) Swap(i, j int) { m.t[i], m.t[j] = m.t[j], m.t[i] }
func (m *MultiSort) Less(i,j int) bool {
	p, q := m.t[i], m.t[j]
	var k int
	for k = 0; k < len(m.compares) - 1; k++ {
		less := m.compares[k]
		switch  {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
		// p == q; try the next comparison.
	}

	return m.compares[k](p, q)
}

func newMultiSort(tracks []*Track, ordering []string) *MultiSort {
	compares := make([]LessFunc, 0)
	for _, v := range ordering {
		if f, ok := sortDict[strings.ToLower(v)]; ok {
			compares = append(compares, f)
		}
	}

	if len(compares) == 0 {
		compares = append(compares, sortByTitle)
	}

	return &MultiSort{t:tracks, compares: compares}
}

func Sort(tracks []*Track, ordering []string) []*Track {
	multiSort := newMultiSort(tracks, ordering)
	sort.Sort(multiSort)
	return multiSort.t
}