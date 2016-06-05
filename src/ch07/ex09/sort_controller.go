package main

import (
	"ch07/ex08/sorting"
	"log"
	"net/http"
	"text/template"
	"time"
)

var musicList = template.Must(template.New("musicList").Parse(`
<table>
<tr style='text-align: left'>
  <th><a href="?Title=true">Title</a></th>
  <th><a href="?Artist=true">Artist</a></th>
  <th><a href="?Album=true">Album</a></th>
  <th><a href="?Year=true">Year</a></th>
  <th><a href="?Length=true">Length</a></th>
</tr>
  {{range .[]*sorting.Track}}
<tr>
  <td>{{.Title}}</td>
  <td>{{.Artist}}</td>
  <td>{{.Album}}</td>
  <td>{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`))

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var tracks = []*sorting.Track{
			{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
			{"Go", "Moby", "Moby", 1992, length("3m37s")},
			{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
			{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
		}
		if err := musicList.Execute(w, tracks); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
