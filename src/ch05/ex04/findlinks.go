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
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	enableType := []string{"a", "script", "img", "link", "tyle", "iframe"}
	enableAttr := []string{"href", "src"}
	// map使っても良いかも。こっちのほうが早い？
	// enableType := map[string]bool{"a": true}
	// if n.Type == html.ElementNode && enableType[n.Data] {
	if n.Type == html.ElementNode && contains(enableType, n.Data) {
		for _, a := range n.Attr {
			if contains(enableAttr, a.Key) {
				links = append(links, a.Val)
			}
		}
	}
	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}
	return links
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
