package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	words, images, err := CountWordsAndImages("http://yahoo.co.jp")
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("words:\t%d\nimages:\t%d\n", words, images)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	return visit(words, images, n)
}

func countWords(word string) (count int) {
	in := bufio.NewScanner(strings.NewReader(word))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		count++
	}
	return
}

func visit(words, images int, n *html.Node) (int, int) {
	reg := regexp.MustCompile("^[ \f\n\r\t\v]*$").Match([]byte(n.Data))
	if n.Type == html.TextNode && !reg && n.Parent.Data != "style" && n.Parent.Data != "script" {
		words += countWords(n.Data)
	}
	enableType := []string{"img"}
	enableAttr := []string{"src"}
	if n.Type == html.ElementNode && contains(enableType, n.Data) {
		for _, a := range n.Attr {
			if contains(enableAttr, a.Key) {
				images++
			}
		}
	}
	if n.FirstChild != nil {
		words, images = visit(words, images, n.FirstChild)
	}
	if n.NextSibling != nil {
		words, images = visit(words, images, n.NextSibling)
	}
	return words, images
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
