package ex11_5

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		input, sep string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		if got := strings.Split(test.input, test.sep); !reflect.DeepEqual(test.want, got) {
			t.Errorf("Split(%q, %q) = %v, want %v", test.input, test.sep, got, test.want)
		}
	}
}
