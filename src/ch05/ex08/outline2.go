package main

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

var depth int

func main() {
	resp, err := http.Get("http://yahoo.co.jp")
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	ElementById(doc, "p")
}

func ElementById(doc *html.Node, id string) *html.Node {
	forEachNode(id, doc, startElement, endElement)
	return doc
}

func forEachNode(id string, n *html.Node, pre func(id string, n *html.Node) bool, post func(n *html.Node) bool) bool {
	if pre != nil {
		if pre(id, n) {
			return true
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if forEachNode(id, c, pre, post) {
			return true
		}
	}

	if post != nil {
		if post(n) {
			return true
		}
	}
	return false
}

func startElement(id string, n *html.Node) bool {
	flag := false
	if n.Type == html.ElementNode {
		var attr string
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				flag = true
			}
			attr += fmt.Sprintf(" %s='%s'", a.Key, a.Val)
		}
		if n.FirstChild == nil {
			fmt.Printf("%*s<%s%s/>\n", depth*2, "", n.Data, attr)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, attr)
			depth++
			return flag
		}
	}
	if n.Type == html.CommentNode {
		fmt.Printf("%*s<!--%s-->\n", depth*2, "", n.Data)
	}
	reg := regexp.MustCompile("^[ \f\n\r\t\v]*$").Match([]byte(n.Data))
	if n.Type == html.TextNode && !reg {
		fmt.Printf("%*s%s\n", depth*2, "", n.Data)
	}
	return flag
}

func endElement(n *html.Node) bool {
	if n.Type == html.ElementNode && n.FirstChild != nil {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
	return false
}
