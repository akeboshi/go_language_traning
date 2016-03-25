package main

import (
	"ch04/ex13/omdb"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		println("Please input movie's title")
		os.Exit(1)
	}
	movie := omdb.SearchMovie(os.Args[1])

	println("Title:\t" + movie.Title)
	println("URL:\t" + movie.Poster)

	omdb.WritePoster(movie.Poster, movie.Title+".png")
}
