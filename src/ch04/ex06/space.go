package main

import "unicode"

func main() {
	runes := []rune("あい　うえお")

	println(string(runes))
	println(string(unicodeSpaceToASCIISpace(runes)))
}

func unicodeSpaceToASCIISpace(str []rune) []rune {
	var result []rune
	for _, r := range str {
		if unicode.IsSpace(r) {
			result = append(result, ' ')
		} else {
			result = append(result, r)
		}
	}
	return result
}
