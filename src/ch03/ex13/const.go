//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import "fmt"

const (
	// TODO: 他にやり方ありそう？
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)
	fmt.Println(PB)
	fmt.Println(EB)
	// TODO: overflowするので苦肉の策
	fmt.Printf("%v000\n", ZB/KB)
	fmt.Printf("%v000000\n", YB/KB/KB)
}
