package main

import (
	"ch04/ex13/omdb"
	"os"
)

func main() {
	movie := omdb.SearchMovie(os.Args[1])
	println(movie.Title)
	println(movie.Poster)

	// omdbに寄付する必要がある？
	// omdb.WritePoster(movie.Poster, movie.Title)
	// omdbから画像がとれないから、とりあえず、ミク画像でも。
	omdb.WritePoster("https://vocadb.net/EntryImg/Artist/488.png", "miku.png")
}
