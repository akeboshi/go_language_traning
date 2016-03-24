package main

import (
	"ch04/ex12/xkcd"
	"flag"
)

func main() {
	cache := flag.Int("c", -1, "指定の番号までcacheするよ")
	search := flag.String("s", "", "search")
	id := flag.Int("i", -1, "get with id")
	flag.Parse()
	if *id != -1 {
		printComic(xkcd.FetchId(*id))
	} else if *search != "" {
		for _, comic := range xkcd.SearchFromIndex(*search) {
			printComic(comic)
		}
	} else if *cache != -1 {
		xkcd.CacheComics(1, *cache)
	}
}

func printComic(comic xkcd.XKCD) {
	println("--------------------------------")
	println("Transcript:\t" + comic.Transcript)
	println("URL:\t" + comic.Img)
	println("--------------------------------")
}
