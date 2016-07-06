// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client struct {
	ch   chan<- string
	name string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				go messageSender(cli.ch, msg)
			}
		case cli := <-entering:
			var list []string
			for c := range clients {
				list = append(list, c.name)
			}
			clients[cli] = true
			if len(list) > 0 {
				str := "member list\n" +
					"-----------\n" +
					strings.Join(list, "\n") +
					"\n-----------"

				go messageSender(cli.ch, str)
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.ch)
		}
	}
}

func messageSender(ch chan<- string, message string) {
	ch <- message
}

func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := "no name"
	input := bufio.NewScanner(conn)
	if input.Scan() {
		who = input.Text()
	}

	ch <- "You are " + who
	messages <- who + " has arrived"
	cli := client{ch, who}
	entering <- cli

	alive := make(chan struct{})
	abort := make(chan struct{})
	go func() {
		for input.Scan() {
			messages <- who + ": " + input.Text()
			alive <- struct{}{}
		}
		abort <- struct{}{}
	}()
	for {
		select {
		case <-time.After(300 * time.Second):
			leaving <- cli
			messages <- who + " has left"
			conn.Close()
			return
		case <-abort:
			leaving <- cli
			messages <- who + " has left"
			conn.Close()
			return
		case <-alive:
		}
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		go func(gmsg string) {
			fmt.Fprintln(conn, gmsg)
		}(msg)
	}
}
