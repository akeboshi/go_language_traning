//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	str := "-1234567890"
	for i := 2; i <= len(str); i++ {
		fmt.Println(comma(str[0:i]))
	}
	for i := len(str); i > 1; i-- {
		fmt.Println(comma(str[1:i]))
	}
}

func comma(word string) string {
	var buf bytes.Buffer

	if strings.HasPrefix(word, "-") {
		buf.WriteByte('-')
		word = word[1:len(word)]
	}

	n := len(word)
	if n < 1 {
		return "N/A"
	}

	for i := 0; i < n; i++ {
		buf.WriteByte(word[i])
		if (n-i-1)%3 == 0 && i+1 != n {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
