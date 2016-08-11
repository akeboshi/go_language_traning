//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package links

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

func Extract(requestURL string) ([]string, error) {
	resp, err := http.Get(requestURL)
	url, _ := url.Parse(requestURL)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", requestURL, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return nil, err
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				fmt.Printf("link: %s	resp: %s\n", link.Host, url.Host)

				if strings.HasSuffix(link.Host, url.Host) {
					save(resp, link)
					if err != nil {
						continue
					}
					links = append(links, link.String())
				}
			}
		}
	}
	forEachNode(doc, visitNode, nil)
	return links, nil
}

func save(resp *http.Response, link *url.URL) {
	respHost := resp.Request.Host
	linkHost := link.Host
	if respHost == linkHost {
		err := writeFile(link)
		if err != nil {
			println(err.Error())
		}
	}
}

func writeFile(link *url.URL) error {
	resp, err := http.Get(link.String())
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return err
	}
	path := link.EscapedPath()
	if path == "" {
		path = "/"
	}
	filename := path[strings.LastIndex(path, "/")+1:]
	if filename == "" || filename == "." {
		filename = "index.html"
	}
	dir := "./" + link.Host + "/" + path[0:strings.LastIndex(path, "/")]
	body = goinReplaceRelativePath(body, link)

	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(dir+"/"+filename, body, os.ModePerm)
	if err != nil {
		return err
	}
	println("saved: " + dir + "/" + filename)
	return nil
}

func goinReplaceRelativePath(body []byte, uri *url.URL) []byte {
	depth := strings.Count(uri.Path, "/")
	rep := regexp.MustCompile(`href="https?://` + uri.Host)
	body = rep.ReplaceAll(body, []byte("href=\""))

	rep = regexp.MustCompile(`href="/`)
	repl := fmt.Sprintf("href=\"%*s", depth, "../")
	body = rep.ReplaceAll(body, []byte(repl))
	return body
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
