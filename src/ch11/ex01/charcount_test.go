// Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestCount(t *testing.T) {
	testDatas := []struct {
		input    string
		expected countData
	}{
		{"foo",
			countData{
				map[rune]int{'f': 1, 'o': 2},
				map[string]int{"letter": 3},
				[utf8.UTFMax + 1]int{0, 3, 0, 0, 0},
				0,
			},
		},
		{"\n\t\t\r",
			countData{
				map[rune]int{'\n': 1, '\t': 2, '\r': 1},
				map[string]int{"control": 4},
				[utf8.UTFMax + 1]int{0, 4, 0, 0, 0},
				0,
			},
		},
		{"🍣ÄあA🍣",
			countData{
				map[rune]int{'Ä': 1, 'あ': 1, 'A': 1, '🍣': 2},
				map[string]int{"letter": 3, "symbol": 2},
				[utf8.UTFMax + 1]int{0, 1, 1, 1, 2},
				0,
			},
		},
		{"🍣A1２\t",
			countData{
				map[rune]int{'A': 1, '1': 1, '２': 1, '🍣': 1, '\t': 1},
				map[string]int{"symbol": 1, "letter": 1, "number": 2, "digit": 2, "control": 1},
				[utf8.UTFMax + 1]int{0, 3, 0, 1, 1},
				0,
			},
		},
	}
	for _, testData := range testDatas {
		buff := bytes.NewBufferString(testData.input)
		data := count(bytes.NewReader(buff.Bytes()))
		checkCountData(data, testData.expected, t)
	}
}

func printErrorCountData(reason string, actual, expected countData, t *testing.T) {
	t.Errorf("%s\n%s\n%s\n\n%s\n%s\n",
		reason,
		`actual data.
===============`,
		countDataForPrint(actual),
		`expected data.
===============`,
		countDataForPrint(expected))
}

func checkCountData(actual, expected countData, t *testing.T) {
	if len(actual.counts) != len(expected.counts) {
		printErrorCountData("counts size is not equal.", actual, expected, t)
	}
	for k, v := range actual.counts {
		if expected.counts[k] != v {
			reason := "counts is not equal.\n"
			reason += fmt.Sprintf("actual.counts[%c] = %d.\nexpected.count[%c] = %d", k, v, k, expected.counts[k])
			printErrorCountData(reason, actual, expected, t)
		}
	}
	if len(actual.types) != len(expected.types) {
		printErrorCountData("types size is not equal.", actual, expected, t)
	}
	for k, v := range actual.types {
		if expected.types[k] != v {
			reason := "types is not equal.\n"
			reason += fmt.Sprintf("actual.types[%s] = %d.\nexpected.type[%s] = %d", k, v, k, expected.types[k])
			printErrorCountData(reason, actual, expected, t)
		}
	}
	for k, v := range actual.utflen {
		if expected.utflen[k] != v {
			reason := "utflen is not equal.\n"
			reason += fmt.Sprintf("actual.utflen[%v] = %v.\nexpected.utflen[%v] = %v", k, v, k, expected.utflen[k])
			printErrorCountData(reason, actual, expected, t)
		}
	}
	if actual.invalid != expected.invalid {
		printErrorCountData("invalid is not equal.", actual, expected, t)
	}
}
