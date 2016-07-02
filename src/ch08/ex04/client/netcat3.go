// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	rAddr, err := net.ResolveTCPAddr("tcp", "localhost:8000")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, rAddr)
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.CloseWrite()
	<-done
	conn.Close()
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
