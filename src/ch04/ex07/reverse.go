package main

func main() {
	str := []byte("あい うえお")
	println(string(str))
	println(string(reverseUTF8(str)))
}

func reverseUTF8(str []byte) []byte {
	reverse(str)
	for i := len(str) - 1; i > 0; {
		charLen := unicodeCharBytes(str[i])
		reverse(str[i-charLen+1 : i+1])
		i -= charLen
	}
	return str
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func unicodeCharBytes(char byte) int {
	if char < 192 {
		return 1
	} else if char < 224 {
		return 2
	} else if char < 240 {
		return 3
	}
	return 4
}
