// Exercis e 3.3: Color each polygon bas ed on its heig ht, so that the peaks are colored red
// ( #ff0000 ) and the val leys blue ( #0000ff ).
package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey;fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if !isFinite(ax) || !isFinite(ay) || !isFinite(bx) || !isFinite(by) ||
				!isFinite(cx) || !isFinite(cy) || !isFinite(dx) || !isFinite(dy) {
				continue
			}
			fmt.Printf("<polygon style='fill: %s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color(i, j), ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Printf("</svg>")
}

func isFinite(num float64) bool {
	return !(math.IsNaN(num) || math.IsInf(num, 0))
}

func corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of ceil (i, j).
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

func color(i, j int) string {
	min := math.NaN()
	max := math.NaN()
	for xoff := 0; xoff <= 1; xoff++ {
		for yoff := 0; yoff <= 1; yoff++ {
			x := xyrange * (float64(i+xoff)/cells - 0.5)
			y := xyrange * (float64(j+yoff)/cells - 0.5)
			z := f(x, y)
			if math.IsNaN(min) || z < min {
				min = z
			}
			if math.IsNaN(max) || z > max {
				max = z
			}
		}
	}

	color := ""
	if math.Abs(max) > math.Abs(min) {
		color = fmt.Sprintf("#%02x0000", 255)
	} else {
		color = fmt.Sprintf("#0000%02x", 255)
	}
	return color
}
