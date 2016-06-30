// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type Tokei struct {
	country string
	host    string
}

type TabWriter struct {
	w       io.Writer
	country string
}

func (t TabWriter) Write(data []byte) (int, error) {
	data_len := len(data)
	_, err := t.w.Write([]byte(t.country + "\t: " + string(data)))
	return data_len, err
}

func main() {
	var tokei []Tokei
	for _, v := range os.Args[1:] {
		tmp := strings.Split(v, "=")
		tokei = append(tokei, Tokei{tmp[0], tmp[1]})
	}

	for _, v := range tokei {
		fmt.Printf("%s %s\n", v.country, v.host)
		go v.handle()
	}

	for {
	}
}

func (t Tokei) handle() {
	conn, err := net.Dial("tcp", t.host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(TabWriter{os.Stdout, t.country}, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
