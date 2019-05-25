package ex11_6

import "testing"

const testVal = 0x1234567890ABCDEF

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(testVal)
	}
}

func BenchmarkPopCount2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount2(testVal)
	}
}

func BenchmarkPopCount3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(testVal)
	}
}
func BenchmarkPopCount4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(testVal)
	}
}
