package palindrome

import "sort"

func isEqual(s sort.Interface, i, j int) bool {
	return !s.Less(i, j) && !s.Less(j, i)
}

func isPalindrome(s sort.Interface) bool {
	lastIdx := s.Len() - 1
	for i := 0; i < s.Len()/2; i++ {
		if !isEqual(s, i, lastIdx-i) {
			return false
		}

	}

	return true
}
