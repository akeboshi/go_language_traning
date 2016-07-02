package main

import (
	"fmt"
	"log"
	"net/url"
	"runtime"
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

func main() {
	uri, err := url.Parse("http://golang.io/")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(uri.Host)
	//log.Println(runtime.NumGoroutine())
	//select {}
}
