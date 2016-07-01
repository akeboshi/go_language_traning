//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch07/ex04/newreader"
	"fmt"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

var depth int

func main() {
	doc, err := newreader.Parse("<html><body>foo bar</body></html>")
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	forEachNode(doc, startElement, endElement)

}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attr string
		for _, a := range n.Attr {
			attr += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attr)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attr)
			depth++
		}
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
	reg := regexp.MustCompile("^[ \f\n\r\t\v]*$").Match([]byte(n.Data))
	if n.Type == html.TextNode && !reg {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
