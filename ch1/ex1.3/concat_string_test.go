package concat_string_test

import (
	"strings"
	"testing"
)

var args = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}

func customConcat(s []string) {
	var str, sep string
	for i := 0; i < len(s); i++ {
		str += sep + s[i]
		sep = " "
	}
}

func BenchmarkCustomConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		customConcat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}
