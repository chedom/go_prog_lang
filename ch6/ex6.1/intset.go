package intset

type IntSet struct {
	words [] uint64
}

func (s *IntSet) Len() int  {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for i:= 0; i < 64; i++ {
			if (word & (1 << uint(i))) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var intSet IntSet
	intSet.words = make([]uint64, 0, cap(s.words))
	copy(intSet.words, s.words)
	return &intSet
}