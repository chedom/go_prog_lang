package word1

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, test := range tests {
		if got := IsPalindrome(test.input); got != test.want {
			t.Errorf("IsPalindrome(%q) = %v", test.input, got)
		}
	}
}


func randomPalindrome(rng *rand.Rand) string {
	max := 25
	n := rng.Intn(max)
	runes := make([]rune, n)
	for i, j := 0, n; i < j; {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			i++
			continue
		}
		runes[j-1] = r
		i++
		j--
	}

	return string(runes)
}

func randomNonPalindrome(rng *rand.Rand) string {
	max := 25
	min := 2
	n := rng.Intn(max - min) + min
	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		for {
			r := rune(rng.Intn(0x1000))
			if unicode.IsLetter(r) {
				runes[i] = r
				break
			}
		}
	}
	return string(runes)
}

func TestRandomPalindromes(t *testing.T) {
	// initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q) = false", p)
		}
	}
}

func TestRandomNonPalindrome(t *testing.T) {
	// initialize a pseudo-random number generator.
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))
	for i := 0; i < 1000; i++ {
		p := randomNonPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindrome(%q)=true", p)
		}
	}
}
