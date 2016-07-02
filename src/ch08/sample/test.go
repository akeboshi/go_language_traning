package main

import (
	"fmt"
	"log"
	"regexp"
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

//func main() {
//log.Println(runtime.NumGoroutine())
//select {}
//}
func main() {
	str := "abcdefabcieajdlabcdeafbe"
	expected := "ABCefABCieajdlABCeafbe"
	rep := regexp.MustCompile(`abcd?`)
	str = rep.ReplaceAllString(str, "ABC")
	if str == expected {
		fmt.Println("OK")
	}
	fmt.Println(str)
}
