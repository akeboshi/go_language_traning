//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fileMap := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		counts := make(map[string]int)
		countLines(os.Stdin, counts)
		fileMap["Stdin"] = counts
	} else {
		for _, filename := range files {
			counts := make(map[string]int)
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			fileMap[filename] = counts
			f.Close()
		}
	}
	for filename, counts := range fileMap {
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%s\t: %d\t%s\n", filename, n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
