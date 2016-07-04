//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func main() {
	ch := make(chan string, len(os.Args))
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	fmt.Println(<-ch)
}

func fetch(url string, ch chan<- string) {
	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	if cancelled() {
		req.Cancel = make(chan struct{})
	}
	done <- struct{}{}

	splitURL := strings.Split(url, "/")
	filename := splitURL[len(splitURL)-1]
	if filename == "" {
		filename = "index.html"
	}
	file, err := os.Create(req.Host + filename)
	if err != nil {
		ch <- fmt.Sprintf("file open err %s: %v", filename, err)
		return
	}
	defer file.Close()

	nbytes, err := io.Copy(file, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	ch <- fmt.Sprintf("%7d %s", nbytes, url)
}
