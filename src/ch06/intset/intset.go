package intset

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint
}

var bitSize = 32 << (^uint(0) >> 63)

func (s *IntSet) Has(x int) bool {
	word, bit := x/bitSize, uint(x%bitSize)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet) Add(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		}
	}
}

func (s *IntSet) Elems() (list []int) {
	for i, word := range s.words {
		for j := 0; j < bitSize; j++ {
			if (word & (1 << uint(j))) != 0 {
				list = append(list, bitSize*i+j)
			}
		}
	}
	return list
}

// (1,1)->0, (1,0)->1 (0,1)->1 (0,0)->0
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] = s.words[i] ^ tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < bitSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", bitSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		if word != 0 {
			for j := 0; j < bitSize; j++ {
				if word&(1<<uint(j)) != 0 {
					count++
				}
			}
		}
	}
	return count
}

func (s *IntSet) Remove(x int) {
	word, bit := x/bitSize, uint(x%bitSize)
	if word >= len(s.words) {
		return
	}
	s.words[word] &^= 1 << bit
}

func (s *IntSet) Clear() {
	s.words = []uint{}
}

func (s *IntSet) Copy() *IntSet {
	cloneWords := make([]uint, len(s.words))
	copy(cloneWords, s.words)
	return &IntSet{cloneWords}
}

func (s *IntSet) AddAll(x ...int) {
	for _, xx := range x {
		s.Add(xx)
	}
}
