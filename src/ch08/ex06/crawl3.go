// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"ch05/ex13/links"
	"flag"
	"fmt"
	"log"
	"os"
)

type wl struct {
	links []string
	depth int
}

type ul struct {
	link  string
	depth int
}

func main() {
	var limitDepth *int = flag.Int("depth", 3, "get with depth")
	flag.Parse()
	worklist := make(chan wl)
	unseenLinks := make(chan ul)

	go func() {
		worklist <- wl{os.Args[1:], 0}
	}()

	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				var foundLinks wl
				foundLinks.links = crawl(link.link)
				foundLinks.depth = link.depth + 1
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		if list.depth < *limitDepth {
			for _, link := range list.links {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- ul{link, list.depth}
				}
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
