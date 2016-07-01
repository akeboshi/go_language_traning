//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sep := ": "
	for index, arg := range os.Args {
		arr := []string{strconv.Itoa(index), arg}
		fmt.Println(strings.Join(arr, sep))
	}
}

/* 出力結果
$ ./echo foo bar
0: ./echo
1: foo
2: bsr
*/
