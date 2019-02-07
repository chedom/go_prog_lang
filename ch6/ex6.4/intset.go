package intset

type IntSet struct {
	words [] uint64
}

func (s *IntSet) Elems() []int {
	res := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j:=0; j < 64; j++ {
			if word & (1<< uint(j)) != 0 {
				res = append(res, i*64+j)
			}
		}
	}
	return res
}
