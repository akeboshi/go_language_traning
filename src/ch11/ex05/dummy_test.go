// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package split

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	data := []struct {
		s        string
		sep      string
		expected int
	}{
		{"a:b:c", ":", 3},
		{"a b c", " ", 3},
		{"", ":", 1},
	}

	for _, d := range data {
		actual := strings.Split(d.s, d.sep)
		if d.expected != len(actual) {
			t.Errorf("Split(%q, %q) returned %d words, want %d", d.s, d.sep, len(actual), d.expected)
		}
	}
}
