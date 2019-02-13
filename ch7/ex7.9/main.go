package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/chedom/go_prog_lang/ch7/ex7.8"
)

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

var tracks = []*multisort.Track{
	{"Go", "Delialah", "Froom the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As i Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}


var trackList = template.Must(template.New("trackList").Parse(`
<h1>Track list, count {{len .}}</h1>
<table border='1'>
	<tr style='text-align:left'>
        <th>
			<a href="?sortBy=title">Title</a>
		</th>
        <th>
 			<a href="?sortBy=artist">Artist</a>
		</th>
        <th>
			<a href="?sortBy=album">Album</a>
		</th>
        <th>
			<a href="?sortBy=year">Year</a>
		</th>
        <th>
			<a href="?sortBy=length">Length</a>
		</th>
	</tr>
{{range .}}
	<tr style='text-align:left'>
    	<th>{{.Title}}</th>
		<th>{{.Artist}}</th>
		<th>{{.Album}}</th>
		<th>{{.Year}}</th>
		<th>{{.Length}}</th>
    </tr>
{{end}}
</table>
`))

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	list := tracks

	if v, ok := r.Form["sortBy"]; ok {
		list = multisort.Sort(list, v)
	}

	if err := trackList.Execute(w, list); err != nil {
		log.Fatal(err)
	}
}
