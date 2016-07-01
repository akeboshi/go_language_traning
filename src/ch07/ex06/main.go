//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch07/ex06/tempconv"
	"flag"
	"fmt"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
