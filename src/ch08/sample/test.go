package main

import (
	"fmt"
	"log"
	"net/url"
	"regexp"
	"runtime"
	"strings"
)

func worker(msg string) <-chan string {
	limit := make(chan int, 5)
	receiver := make(chan string)
	go func() {
		for i := 0; i < 100; i++ {
			log.Println(runtime.NumGoroutine())
			limit <- 1
			go func(i int) {
				msg := fmt.Sprintf("%d %s done", i, msg)
				receiver <- msg
				<-limit
			}(i)
		}
	}()
	return receiver
}

//func main() {
//log.Println(runtime.NumGoroutine())
//select {}
//}
func main() {
	here := "http://aruga.src.jp/foo"
	hereUrl, _ := url.Parse(here)
	depth := strings.Count(hereUrl.Path, "/")
	urls := []string{
		"href=\"/foo/bar/hoge",                // ../foo/bar/hoge or ./bar/hoge
		"href=\"./hoge",                       // ./hoge
		"href=\"http://aruga.src.jp/bar",      // ../bar
		"href=\"https://aruga.src.jp/foo/bar", // ../foo/bar or ./bar
	}
	body := ""
	for _, url := range urls {
		body += url + "\n"
	}
	rep := regexp.MustCompile(`href="https?://` + hereUrl.Host)
	body = rep.ReplaceAllString(body, "href=\"")

	rep = regexp.MustCompile(`href="/`)
	repl := fmt.Sprintf("href=\"%*s", depth, "../")
	body = rep.ReplaceAllString(body, repl)
	println(body)
}
