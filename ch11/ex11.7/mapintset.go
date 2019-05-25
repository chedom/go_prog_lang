package intset

import (
	"bytes"
	"fmt"
	"sort"
)

type MapIntSet map[int]struct{}

func NewMapIntSet() IntSet {
	return MapIntSet(make(map[int]struct{}, 0))
}

func (m MapIntSet) Has(x int) bool {
	_, ok := m[x]
	return ok
}

func (m MapIntSet) Add(x int) {
	m[x] = struct{}{}
}

func (m MapIntSet) UnionWith(other IntSet) {
	if m2, ok := other.(MapIntSet); ok {
		for k := range m2 {
			m[k] = struct{}{}
		}
	}
}

func (m MapIntSet) String() string {
	sl := make([]int, 0, len(m))
	for k := range m {
		sl = append(sl, k)
	}

	sort.Ints(sl)

	b := &bytes.Buffer{}
	b.WriteByte('{')

	for i, v := range sl {
		if i != 0 {
			b.WriteByte(' ')
		}
		fmt.Fprintf(b, "%d", v)
	}

	b.WriteByte('}')

	return b.String()
}

func (m MapIntSet) Len() int {
	return len(m)
}

func (m MapIntSet) Remove(x int) {
	delete(m, x)
}

func (m MapIntSet) Clear() {
	for k := range m {
		delete(m, k)
	}
}