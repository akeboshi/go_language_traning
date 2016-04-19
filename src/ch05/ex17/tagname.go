package main

import (
	"fmt"
	"net/http"

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
	images := ElementsByTagName(doc, "img")
	println(len(images))
	for _, img := range images {
		for _, a := range img.Attr {
			if a.Key == "src" {
				println(a.Val)
			}
		}
	}
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	println(len(headings))
}

func ElementsByTagName(doc *html.Node, name ...string) (list []*html.Node) {
	var forEachNode func(*html.Node)
	forEachNode = func(n *html.Node) {
		if n.Type == html.ElementNode && contain(n.Data, name) {
			list = append(list, n)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			forEachNode(c)
		}
	}
	forEachNode(doc)
	return
}

func contain(str string, list []string) bool {
	for _, l := range list {
		if l == str {
			return true
		}
	}
	return false
}
