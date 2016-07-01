//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	mux.Handle("/create", http.HandlerFunc(db.create))
	mux.Handle("/delete", http.HandlerFunc(db.delete))
	mux.Handle("/update", http.HandlerFunc(db.update))
	mux.Handle("/list", http.HandlerFunc(db.list))
	mux.Handle("/price", http.HandlerFunc(db.price))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type dollars float32
type database map[string]dollars

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

// POST であるべき
func (db database) create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "あいてむとねだんをしていしてね")
	}
	_, ok := db[item]
	if ok {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "していされたあいてむはもうそんざいしてるよ")
	}
	intPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ねだんはすうじでしていしてね")
	}
	db[item] = dollars(intPrice)
	fmt.Fprintf(w, "%s を %s でとうろくしたよ！", item, db[item])
}

// DELETE
func (db database) delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "さくじょするあいてむをえらんでね")
		return
	}
	_, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s はとうろくされてないよ！\n", item)
		return
	}
	delete(db, item)
	fmt.Fprintf(w, "%s のさくじょにせいこう！", item)
}

// PUT
func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price := req.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "あいてむとねだんをしていしてね")
		return
	}
	intPrice, err := strconv.ParseFloat(price, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "ねだんはすうじでしていしてね")
		return
	}
	db[item] = dollars(intPrice)
	fmt.Fprintf(w, "%s を %s でこうしんしたよ！", item, db[item])
}

var itemList = template.Must(template.New("itemList").Parse(`
<table>
<tr>
  <th><strong>Item</strong></th>
  <th><strong>Price</strong></th>
</tr>
{{range $item, $price := .}}
<tr>
<td>{{ $item }}</td>
<td>{{ $price }}</td>
</tr>
{{end}}
</table>
`))

func (db database) list(w http.ResponseWriter, req *http.Request) {
	if err := itemList.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "%s はとうろくされてないよ！\n", item)
		return
	}
	fmt.Fprintf(w, "%s\n", price)
}
