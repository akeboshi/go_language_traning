// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup
	text := make(chan string)
	go func() {
		for input.Scan() {
			text <- input.Text()
		}
	}()
	var timeOut sync.WaitGroup
	done := make(chan struct{})
	go func() {
		timeOut.Add(1)
		go func() {
			timeOut.Wait()
			done <- struct{}{}
		}()
		<-time.After(10 * time.Second)
		timeOut.Done()
	}()

	for {
		select {
		case x := <-text:
			wg.Add(1)
			go echo(c, x, 1*time.Second, &wg)
			go func() {
				timeOut.Add(1)
				<-time.After(10 * time.Second)
				timeOut.Done()
			}()

		case <-done:
			wg.Wait()
			c.Close()
		}
	}
}
