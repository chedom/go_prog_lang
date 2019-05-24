package intset

import (
	"testing"
)

func newIntSets() []IntSet {
	return []IntSet{NewMapIntSet(), &BitIntSet{}}
}

func TestLenZeroInitially(t *testing.T) {
	for _, s := range newIntSets() {
		len := s.Len()
		if len != 0 {
			t.Errorf("%T.Len(): got %d, want 0", s, len)
		}
	}
}

func TestAdd(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(1)
		s.Add(3)

		if s.Len() != 2 {
			t.Errorf("%T.Len(): got %d, want 2", s, s.Len())
		}

		if !s.Has(1) {
			t.Errorf("%T.Has(%d): got %t, want %t", s, 1, s.Has(1), !s.Has(1))
		}

		if !s.Has(3) {
			t.Errorf("%T.Has(%d): got %t, want %t", s, 1, s.Has(3), !s.Has(3))
		}
	}
}

func TestRemove(t *testing.T) {
	for _, s := range newIntSets() {
		s.Add(0)
		s.Remove(0)
		if s.Has(0) {
			t.Errorf("%T: want zero removed, got %s", s, s)
		}
	}
}