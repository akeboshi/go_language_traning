// Copyright (c) 2016 by akeboshi. All Rights Reserved.

package main

import (
	"ch10/ex02/myarchive"
	_ "ch10/ex02/myarchive/tar"
	_ "ch10/ex02/myarchive/zip"
	"fmt"
	"os"
)

func main() {
	kind, err := myarchive.Extract(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
	println("file type is " + kind)
}
