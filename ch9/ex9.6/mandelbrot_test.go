package mandelbrot

import (
	"io/ioutil"
	"testing"
)

func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createImg(ioutil.Discard)
	}
}
