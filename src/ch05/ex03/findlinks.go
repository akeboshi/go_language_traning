package main

import (
	"fmt"
	"os"
	"regexp"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, v := range visit(nil, doc) {
		fmt.Printf("%s\n", v)
	}
}

func visit(contents []string, n *html.Node) []string {
	reg := regexp.MustCompile("^[ \f\n\r\t\v]*$").Match([]byte(n.Data))
	if n.Type == html.TextNode && !reg && n.Parent.Data != "style" && n.Parent.Data != "script" {
		contents = append(contents, n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		contents = visit(contents, c)
	}
	return contents
}
