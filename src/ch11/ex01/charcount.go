//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type countData struct {
	counts  map[rune]int
	types   map[string]int
	utflen  [utf8.UTFMax + 1]int
	invalid int
}

func newCountData() countData {
	var data countData
	data.counts = make(map[rune]int)
	data.types = make(map[string]int)
	return data
}

func count(r io.Reader) countData {
	data := newCountData()
	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			data.invalid++
			continue
		}
		data.counts[r]++
		data.utflen[n]++

		if unicode.IsLetter(r) {
			data.types["letter"]++
		}
		if unicode.IsControl(r) {
			data.types["control"]++
		}
		if unicode.IsDigit(r) {
			data.types["digit"]++
		}
		if unicode.IsNumber(r) {
			data.types["number"]++
		}
		if unicode.IsSymbol(r) {
			data.types["symbol"]++
		}
	}
	return data
}

func printCountData(data countData) {
	fmt.Printf(countDataForPrint(data))
}

func countDataForPrint(data countData) string {
	var str string
	str += fmt.Sprintf("rune\tcount\n")
	for c, n := range data.counts {
		str += fmt.Sprintf("%q\t%d\n", c, n)
	}
	str += fmt.Sprint("\nlen\tcount\n")
	for i, n := range data.utflen {
		if i > 0 {
			str += fmt.Sprintf("%d\t%d\n", i, n)
		}
	}

	str += fmt.Sprint("\ntype\tcount\n")
	for k, v := range data.types {
		str += fmt.Sprintf("%s\t%d\n", k, v)
	}
	if data.invalid > 0 {
		str += fmt.Sprintf("\n%d invalid UTF-8 characters\n", data.invalid)
	}
	return str
}

func main() {
	data := count(os.Stdin)
	printCountData(data)
}
