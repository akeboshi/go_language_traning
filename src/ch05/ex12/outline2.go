package main

import (
	"fmt"
	"net/http"
	"regexp"

	"golang.org/x/net/html"
)

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
	var depth int
	forEachNode(doc, func(n *html.Node) {
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
	}, func(n *html.Node) {
		if n.Type == html.ElementNode && n.FirstChild != nil {
			depth--
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	})
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
