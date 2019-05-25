package intset

import (
	"math/rand"
	"testing"
)

func newIntSets() []IntSet {
	return []IntSet{NewMapIntSet(), NewBitIntSet(), NewBitInt32Set()}
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


// General benchmark function
const max = 32000

func addRandom(set IntSet, n int) {
	for i:= 0; i < n; i++ {
		set.Add(rand.Intn(max))
	}
}

func benchHas(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.Has(rand.Intn(max))
	}
}

func benchAdd(b *testing.B, set IntSet, n int) {
	for i:= 0; i < b.N; i++ {
		for j := 0; j < n; j++ {
			set.Add(rand.Intn(max))
		}
		set.Clear()
	}
}

func benchUnionWith(bm *testing.B, a, b IntSet, n int) {
	addRandom(a, n)
	addRandom(b, n)
	for i := 0; i < bm.N; i++ {
		a.UnionWith(b)
	}
}

func benchString(b *testing.B, set IntSet, n int) {
	addRandom(set, n)
	for i := 0; i < b.N; i++ {
		set.String()
	}
}

// benchmark MapIntSet

func BenchmarkMapIntSetAdd10(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 10)
}

func BenchmarkMapIntSetAdd100(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 100)
}

func BenchmarkMapIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetHas10(b *testing.B) {
	benchHas(b, NewMapIntSet(), 10)
}


func BenchmarkMapIntSetHas100(b *testing.B) {
	benchHas(b, NewMapIntSet(), 100)
}


func BenchmarkMapIntSetHas1000(b *testing.B) {
	benchHas(b, NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 10)
}


func BenchmarkMapIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 100)
}


func BenchmarkMapIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewMapIntSet(), NewMapIntSet(), 1000)
}

func BenchmarkMapIntSetString10(b *testing.B) {
	benchString(b, NewMapIntSet(), 10)
}


func BenchmarkMapIntSetString100(b *testing.B) {
	benchString(b, NewMapIntSet(), 100)
}


func BenchmarkMapIntSetString1000(b *testing.B) {
	benchString(b, NewMapIntSet(), 1000)
}


// benchmark bitIntSet32

func BenchmarkBitIntSet32Add10(b *testing.B) {
	benchAdd(b, NewBitInt32Set(), 10)
}

func BenchmarkBitIntSet32Add100(b *testing.B) {
	benchAdd(b, NewBitInt32Set(), 100)
}

func BenchmarkBitIntSet32Add1000(b *testing.B) {
	benchAdd(b, NewBitInt32Set(), 1000)
}

func BenchmarkBitIntSet32Has10(b *testing.B) {
	benchHas(b, NewBitInt32Set(), 10)
}


func BenchmarkBitIntSet32Has100(b *testing.B) {
	benchHas(b, NewBitInt32Set(), 100)
}


func BenchmarkBitIntSet32Has1000(b *testing.B) {
	benchHas(b, NewBitIntSet(), 1000)
}

func BenchmarkBitInt32SetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewBitInt32Set(), NewBitInt32Set(), 10)
}


func BenchmarkBitInt32SetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewBitInt32Set(), NewBitInt32Set(), 100)
}


func BenchmarkBitInt32SetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewBitInt32Set(), NewBitInt32Set(), 1000)
}

func BenchmarkBitInt32SetString10(b *testing.B) {
	benchString(b, NewBitInt32Set(), 10)
}


func BenchmarkBitInt32SetString100(b *testing.B) {
	benchString(b, NewBitInt32Set(), 100)
}


func BenchmarkBitInt32SetString1000(b *testing.B) {
	benchString(b, NewBitInt32Set(), 1000)
}

// benchmark bitIntSet

func BenchmarkBitIntSetAdd10(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 10)
}

func BenchmarkBitIntSetAdd100(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 100)
}

func BenchmarkBitIntSetAdd1000(b *testing.B) {
	benchAdd(b, NewBitIntSet(), 1000)
}

func BenchmarkBitIntSetHas10(b *testing.B) {
	benchHas(b, NewBitIntSet(), 10)
}


func BenchmarkBitIntSetHas100(b *testing.B) {
	benchHas(b, NewBitIntSet(), 100)
}


func BenchmarkBitIntSetHas1000(b *testing.B) {
	benchHas(b, NewBitIntSet(), 1000)
}

func BenchmarkBitIntSetUnionWith10(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 10)
}


func BenchmarkBitIntSetUnionWith100(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 100)
}


func BenchmarkBitIntSetUnionWith1000(b *testing.B) {
	benchUnionWith(b, NewBitIntSet(), NewBitIntSet(), 1000)
}

func BenchmarkBitIntSetString10(b *testing.B) {
	benchString(b, NewBitIntSet(), 10)
}


func BenchmarkBitIntSetString100(b *testing.B) {
	benchString(b, NewBitIntSet(), 100)
}


func BenchmarkBitIntSetString1000(b *testing.B) {
	benchString(b, NewBitIntSet(), 1000)
}