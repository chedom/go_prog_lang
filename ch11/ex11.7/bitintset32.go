package intset

import (
	"bytes"
	"fmt"
)

type BitInt32Set struct {
	words []uint32
}

func NewBitInt32Set() *BitInt32Set {
	return &BitInt32Set{}
}

func (s *BitInt32Set) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *BitInt32Set) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *BitInt32Set) UnionWith(t IntSet) {
	if t2, ok := t.(*BitInt32Set); ok {
		for i, tword := range t2.words {
			if i < len(s.words) {
				s.words[i] |= tword
			} else {
				s.words = append(s.words, tword)
			}
		}

	}
}

func (s *BitInt32Set) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func popcount32(x uint32) int {
	count := 0
	for x != 0 {
		count++
		x &= x - 1
	}
	return count
}

// return the number of elements
func (s *BitInt32Set) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount32(word)
	}
	return count
}

// remove x from the set
func (s *BitInt32Set) Remove(x int) {
	word, bit := x/32, uint(x%32)
	s.words[word] &^= 1 << bit
}

func (s *BitInt32Set) Clear()  {
	for i := range s.words {
		s.words[i] = 0
	}
}
