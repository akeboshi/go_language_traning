// Copyright (c) 2016 by akeboshi. All Rights Reserved.
/*
  標準入力から入力された zip/tar ファイルを解凍する
*/
package myarchive

import (
	"fmt"
	"io/ioutil"
	"os"
)

type detectMethod func([]byte) bool
type extractMethod func([]byte, string) error

var detectContentTypeMethods = make(map[string]detectMethod)
var extractMethods = make(map[string]extractMethod)

var unknown string = "unknown"

func RegisterFormat(kind string, detect detectMethod, extract extractMethod) {
	if kind == unknown {
		panic("can't register unknown type")
	}
	detectContentTypeMethods[kind] = detect
	extractMethods[kind] = extract
}

func detectContentType(b []byte) string {
	for kind, method := range detectContentTypeMethods {
		if method(b) {
			return kind
		}
	}
	return unknown
}

func Extract(file *os.File) (string, error) {
	b, _ := ioutil.ReadAll(file)

	kind := detectContentType(b)
	if kind == unknown {
		return unknown, fmt.Errorf("can't detect content type.\n")
	}
	// extract to current directory
	err := extractMethods[kind](b, "./")
	fmt.Println("done.")
	return kind, err
}
