//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	printHashs(getOption())
}

func getOption() (string, []func(string)) {
	var hashFlag [3]*bool
	hashFlag[0] = flag.Bool("l", false, "use sha256")
	hashFlag[1] = flag.Bool("m", false, "use sha384")
	hashFlag[2] = flag.Bool("h", false, "use sha512")
	flag.Parse()
	word := flag.Arg(0)
	if word == "" {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "フラグの後に文字列いれてほしい")
		os.Exit(1)
	}

	var useFunctions []func(string)
	functions := []func(string){printSHA256, printSHA384, printSHA512}
	for i, f := range hashFlag {
		if *f {
			useFunctions = append(useFunctions, functions[i])
		}
	}

	if len(useFunctions) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	return word, useFunctions
}

func printHashs(word string, functions []func(string)) {
	for _, function := range functions {
		function(word)
	}
}

func printSHA256(word string) {
	fmt.Printf("SHA256:%x\n", sha256.Sum256([]byte(word)))
}
func printSHA384(word string) {
	fmt.Printf("SHA384:%x\n", sha512.Sum384([]byte(word)))
}
func printSHA512(word string) {
	fmt.Printf("SHA512:%x\n", sha512.Sum512([]byte(word)))
}
