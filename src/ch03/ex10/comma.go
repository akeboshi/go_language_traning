//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "1234567890"
	for i := 0; i < len(str); i++ {
		fmt.Println(comma(str[0:i]))
	}
}

func comma(word string) string {
	var buf bytes.Buffer
	n := len(word)
	for i := 0; i < n; i++ {
		buf.WriteByte(word[i])
		if (n-i-1)%3 == 0 && i+1 != n {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}
