// Exercis e 3.4: Fo llowing the appro ach of the Lissajous example in Sec tion 1.7, cons truct a web
// server that computes sur faces and writes SVG dat a to the client. The ser ver must set the Con-
// tent-Type he ader like this:
package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0
	angle   = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", svgHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func svgHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	width := float64(600)
	height := float64(320)
	if val := r.FormValue("width"); val != "" {
		if n, err := strconv.ParseFloat(val, 64); err != nil {
			width = n
		} else {
			fmt.Fprintln(os.Stderr, "Error while parsing number: %v\n", err)
			os.Exit(1)
		}
	}
	if val := r.FormValue("height"); val != "" {
		if n, err := strconv.ParseFloat(val, 64); err != nil {
			height = n
		} else {
			fmt.Fprintln(os.Stderr, "Error while parsing number: %v\n", err)
			os.Exit(1)
		}
	}
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey;fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, width, height)
			bx, by := corner(i, j, width, height)
			cx, cy := corner(i, j+1, width, height)
			dx, dy := corner(i+1, j+1, width, height)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func corner(i, j int, width, height float64) (float64, float64) {
	// Find point (x, y) at corner of ceil (i, j).
	xyscale := width / 2 / xyrange
	zscale := height * 0.4
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x, y, z) isometrically onto 2-D SVF canvas (sx, sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
