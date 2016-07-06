// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"fmt"
	"time"
)

var a chan struct{} = make(chan struct{})
var b chan struct{} = make(chan struct{})
var result chan int = make(chan int)

func main() {
	go pong()
	go ping()
	fmt.Println(<-result)
}

func ping() {
	var count int
	b <- struct{}{}
	t := time.After(1 * time.Second)
	for {
		select {
		case <-t:
			result <- count
		default:
			count++
			b <- <-a
		}
	}
}

func pong() {
	for {
		a <- <-b
	}
}
