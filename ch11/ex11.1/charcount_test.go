package ex11_1

import (
	"reflect"
	"strings"
	"testing"
)

type testItem struct {
	input string

	counts  map[rune]int
	utflen  map[int]int
	invalid int
}

func TestCharcount(t *testing.T) {
	tests := []testItem{
		{
			input:   "Hi, 世.",
			counts:  map[rune]int{'H': 1, 'i': 1, ',': 1, ' ': 1, '世': 1, '.': 1},
			utflen:  map[int]int{1: 5, 3: 1},
			invalid: 0,
		},
	}

	for _, test := range tests {
		gotCounts, gotUtflen, gotInvalid := Charcount(strings.NewReader(test.input))

		if !reflect.DeepEqual(gotCounts, test.counts) {
			t.Errorf("%q counts: got %v, want %v", test.input, gotCounts, test.counts)
		}

		if !reflect.DeepEqual(gotUtflen, test.utflen) {
			t.Errorf("%q utflen: got %v, want %v", test.input, gotUtflen, test.utflen)
		}

		if !reflect.DeepEqual(gotCounts, test.counts) {
			t.Errorf("%q invalid: got %v, want %v", test.input, gotInvalid, test.invalid)
		}
	}
}
