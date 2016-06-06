package main

import (
	"ch07/ex08/sorting"
	"log"
	"net/http"
	"sort"
	"strings"
	"text/template"
	"time"
)

var musicList = template.Must(template.New("musicList").Funcs(template.FuncMap{"setSortList": setSortList}).Parse(`
<table>
<tr style='text-align: left'>
  <th><a href='?sort={{setSortList "Title" .SortList}}'>Title</a></th>
  <th><a href='?sort={{setSortList "Artist" .SortList}}'>Artist</a></th>
  <th><a href='?sort={{setSortList "Album" .SortList}}'>Album</a></th>
  <th><a href='?sort={{setSortList "Year" .SortList}}'>Year</a></th>
  <th><a href='?sort={{setSortList "Length" .SortList}}'>Length</a></th>
</tr>
{{range .T}}
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
		var sortList = []string{}
		if r.URL.Query().Get("sort") != "" {
			sortList = strings.Split(r.URL.Query().Get("sort"), ",")
		}
		result := sorting.CustomSort{
			tracks,
			sortList,
		}
		sort.Sort(result)
		if err := musicList.Execute(w, result); err != nil {
			log.Fatal(err)
		}
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func test(a string) string {
	return "val"
}

func setSortList(sort string, list []string) string {
	var newList = []string{sort}
	for _, l := range list {
		if sort != l {
			newList = append(newList, l)
		}
	}
	return strings.Join(newList, ",")
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
