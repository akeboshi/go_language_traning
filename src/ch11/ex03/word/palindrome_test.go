// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25)
	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		rn := rng.Intn(0x1100)
		r := rune(rn)
		if rn > 0x1000 && rn < 0x1050 {
			runes[i] = ' '
			n--
		} else if rn > 0x1050 && rn < 0x1100 {
			runes[i] = ','
			n--
		} else {
			runes[i] = r
			runes[n-1-i] = r
		}
	}
	return string(runes)
}

func randomNotPalindrome(rng *rand.Rand) string {
	n := rng.Intn(23) + 2
	runes := make([]rune, n)
	var letters []*rune
	for i := 0; i < n; i++ {
		r := rune(rng.Intn(0x1000))
		runes[i] = r
		if unicode.IsLetter(r) {
			letters = append(letters, &r)
		}
	}
	if len(letters) < 2 {
		return "ab" // not palindrome
	}
	if unicode.ToLower(*letters[0]) == unicode.ToLower(*letters[len(letters)-1]) {
		if unicode.ToLower(*letters[0]) == 'a' {
			*letters[0] = 'b'
		} else {
			*letters[0] = 'a'
		}
	}

	return string(runes)
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("isNotPalindrome(%q) = true", p)
		}
	}
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("isPalindrome(%q) = false", p)
		}
	}
}
