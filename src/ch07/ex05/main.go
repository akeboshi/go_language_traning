//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch07/ex04/newreader"
	"ch07/ex05/limitreader"
	"fmt"
	"io/ioutil"
)

func main() {
	reader := newreader.NewReader("foo bar")
	reader = limitreader.LimitReader(reader, 3)
	foo, _ := ioutil.ReadAll(reader)

	fmt.Println(string(foo))

}
