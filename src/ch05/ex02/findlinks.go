//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for k, v := range visit(map[string]int{}, doc) {
		fmt.Printf("%s\t: %2d\n", k, v)
	}
}

func visit(types map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		types[n.Data]++
	}
	if n.FirstChild != nil {
		types = visit(types, n.FirstChild)
	}
	if n.NextSibling != nil {
		types = visit(types, n.NextSibling)
	}
	return types
}
