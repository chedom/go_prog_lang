// Mandelbrot emits a PNG image of the Mandelbrot fractial.
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
	"os"
	"sync"
)

func main() {
	createImg(os.Stdout)
}

func createImg(out io.Writer) {
	type item struct {
		px    int
		py    int
		color color.Color
	}

	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	result := make(chan item)
	var wg sync.WaitGroup

	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			wg.Add(1)

			go func(px, py int, y float64) {
				defer wg.Done()

				x := float64(px)/width*(xmax-xmin) + xmin
				z := complex(x, y)
				result <- item{px, py, mandelbrot(z)}
			}(px, py, y)

		}
	}

	go func() {
		wg.Wait()
		close(result)
	}()

	for v := range result {
		img.Set(v.px, v.py, v.color)
	}

	png.Encode(out, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
