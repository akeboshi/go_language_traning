//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import "fmt"

func main() {
	a := []string{"foo", "bar", "bar", "bar", "bar", "huga", "hoge", "hoge"}
	fmt.Println(removeSameStr(a))
}

func removeSameStr(words []string) []string {
	for i := 1; i < len(words); i++ {
		if words[i] == words[i-1] {
			words = remove(words, i)
			i--
		}
	}
	return words
}

func remove(slice []string, i int) []string {
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}
