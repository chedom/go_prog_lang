package ex12_9

import (
	"fmt"
	"io"
	"strconv"
	"text/scanner"
)

type Token interface{}

type Symbol string
type String string

type Int int

type StartList struct{}

type EndList struct{}

type Decoder struct {
	scan  scanner.Scanner
	depth int
}

func NewDecoder(r io.Reader) *Decoder {
	scan := scanner.Scanner{Mode: scanner.GoTokens}
	scan.Init(r)
	return &Decoder{scan, 0}
}

func (d *Decoder) Token() (Token, error) {
	t := d.scan.Scan()

	switch t {
	case scanner.EOF:
		return nil, io.EOF

	case scanner.Ident:
		return Symbol(d.scan.TokenText()), nil

	case scanner.String:
		v, err := strconv.Unquote(d.scan.TokenText())
		if err != nil {
			return nil, err
		}
		return String(v), nil

	case scanner.Int:
		n, err := strconv.ParseInt(d.scan.TokenText(), 10, 32)
		if err != nil {
			return nil, err
		}
		return Int(n), nil

	case '(':
		d.depth++
		return StartList{}, nil

	case ')':
		d.depth--
		return EndList{}, nil

	default:
		pos := d.scan.Pos()
		return nil, fmt.Errorf("unexpected token %s at L%d:C%d", scanner.TokenString(t), pos.Line, pos.Column)
	}
}
