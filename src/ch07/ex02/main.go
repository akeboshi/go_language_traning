//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch07/ex02/countingwriter"
	"fmt"
	"os"
)

func main() {
	writer, count := countingwriter.CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "aaa\n")
	println(*count)
	fmt.Fprintf(writer, "bar\n")
	println(*count)

}
