// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	var port *int = flag.Int("port", 8000, "get with port")
	flag.Parse()
	host := fmt.Sprintf("localhost:%d", *port)
	println(host)
	listener, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Listen by port %d", *port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
