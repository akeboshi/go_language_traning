// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"bufio"
	"fmt"
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
	fmt.Printf("user name:")
	user, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		user = "no name"
	}

	conn, err := net.DialTCP("tcp", nil, rAddr)

	sendUserName(conn, user)
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

func sendUserName(dst io.Writer, name string) {
	w := bufio.NewWriter(dst)
	w.WriteString(name)
	w.Flush()
}

func mustCopy(dst io.Writer, src io.Reader) {
	r := bufio.NewReader(src)
	w := bufio.NewWriter(dst)
	for {
		s, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		w.WriteString(s)
		w.Flush()
	}
}
