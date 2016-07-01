//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch02/ex03/popcount"
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x diff %x = %d\n", c1, c2, popDifCount(c1, c2))
}

func popDifCount(x, y [32]byte) int {
	var count int
	for i := 0; i < 32; i++ {
		count += popcount.PopCountHackers(uint64(x[i] ^ y[i]))
	}
	return count
}
