package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	local, n, err := fetch("https://golang.org/")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	fmt.Printf("%s: %d byte\n", local, n)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		if err == nil {
			err = f.Close()
		}
	}()
	n, err = io.Copy(f, resp.Body)

	return local, n, err
}
