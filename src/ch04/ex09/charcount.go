//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	file, err := os.Open("alice.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	in := bufio.NewScanner(file)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		counts[in.Text()]++
	}

	fmt.Println("count\tword")
	for word, count := range counts {
		fmt.Printf("%d\t%s\n", count, word)
	}
}
