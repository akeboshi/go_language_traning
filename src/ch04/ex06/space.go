//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"unicode"
	"unicode/utf8"
)

func main() {
	str := []byte("あい　うえお")
	println(string(str))
	println(string(unicodeSpaceToASCIISpace(str)))
}

func unicodeSpaceToASCIISpace(str []byte) []byte {
	var result []byte
	for i := 0; i < len(str); i++ {
		charLen := unicodeCharBytes(str[i])
		bStr := str[i : i+charLen]
		i += charLen - 1
		uStr, _ := utf8.DecodeRune(bStr)
		if unicode.IsSpace(uStr) {
			result = append(result, ' ')
		} else {
			result = append(result, bStr...)
		}
	}
	return result
}

func unicodeCharBytes(char byte) int {
	if char < 192 {
		return 1
	} else if char < 224 {
		return 2
	} else if char < 240 {
		return 3
	} else {
		return 4
	}
}
