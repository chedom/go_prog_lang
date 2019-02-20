package eval

import (
	"fmt"
	"strings"
)

func (v Var) String() string {
	return string(v)
}

func (l literal) String() string {
	return fmt.Sprintf("%g", l)
}

func (u unary) String() string {
	return string(u.op) + u.x.String()
}

func (b binary) String() string {
	return fmt.Sprintf(
		"(%s %s %s)",
		b.x.String(),
		string(b.op),
		b.y.String(),
	)

}

func (c call) String() string {
	args := make([]string, 0, len(c.args))
	for _, a := range c.args {
		args = append(args, a.String())
	}
	return c.fn + "("+ strings.Join(args, ", ") + ")"
}

