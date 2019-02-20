package eval

import (
	"testing"
)

func TestString(t *testing.T) {
	tests := []string{
		"-1 + -x",
		"-1 - x",
		"5 / 9 * (F - 32)",
		"pow(x, 3) + pow(y, 3)",
		"sqrt(A / pi)",
	}

	for _, i := range tests {
		expr, err := Parse(i)
		if err != nil {
			t.Error(err)
			continue
		}
		expr2, err := Parse(expr.String())
		if err != nil {
			t.Error(err)
		}
		if expr.String() != expr2.String() {
			t.Errorf("%s != %s", expr.String(), expr2.String())
		}
	}
}
