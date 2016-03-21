package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "anagram"
	str2 := "aaangrm"
	nonStr := "ahagram"
	fmt.Println(anagram(str, str2))
	fmt.Println(anagram(str, nonStr))
}

func anagramByte(str1, str2 string) bool {
	strByte1 := []byte(str1)
	strByte2 := []byte(str2)
	n1 := len(str1)
	n2 := len(str2)
	if n1 != n2 {
		return false
	}
	for i := 0; i < n1; i++ {
		s := []byte{strByte1[i]}
		if bytes.Count(strByte1, s) != bytes.Count(strByte2, s) {
			return false
		}
	}
	return true
}

func anagram(str1, str2 string) bool {
	n1 := len(str1)
	n2 := len(str2)
	if n1 != n2 {
		return false
	}
	for i := 0; i < n1; i++ {
		s := str1[i : i+1]
		if strings.Count(str1, s) != strings.Count(str2, s) {
			return false
		}
	}
	return true
}
