package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*Track{
	{"Go", "Delialah", "Froom the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As i Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "-----", "-----", "-----", "-----")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush()
}

// sort by artist
type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// sort by year
type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// custom sort
type customSort struct {
	t    []*Track
	less func(i, j *Track) bool
}

func (c customSort) Len() int           { return len(c.t) }
func (c customSort) Less(i, j int) bool { return c.less(c.t[i], c.t[j]) }
func (c customSort) Swap(i, j int)      { c.t[i], c.t[j] = c.t[j], c.t[i] }

func main() {
	printTracks(tracks)
	fmt.Println()

	sort.Sort(byArtist(tracks))
	printTracks(tracks)
	fmt.Println()

	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)
	fmt.Println()

	sort.Sort(byYear(tracks))
	printTracks(tracks)
	fmt.Println()

	sort.Sort(customSort{tracks, func(i, j *Track) bool {
		if i.Title != j.Title {
			return i.Title < j.Title
		}
		if i.Year != j.Year {
			return i.Year < j.Year
		}
		if i.Length < j.Length {
			return i.Length < j.Length
		}
		return false
	}})
	printTracks(tracks)
	fmt.Println()
}
