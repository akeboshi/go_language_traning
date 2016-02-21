package main

import (
	"ch02/ex03/popcount"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	num := rand.Uint32()
	result := measure(popcount.PopCount, uint64(num))
	fmt.Println("Popcount:        " + result)
	result = measure(popcount.PopCountLoopVer, uint64(num))
	fmt.Println("PopcountLoop:    " + result)
	result = measure(popcount.PopCountCheck64, uint64(num))
	fmt.Println("PopcountCheck64: " + result)
}

func measure(f func(uint64) int, num uint64) string {
	start := time.Now()
	for i := 1; i < 1000000; i++ {
		_ = f(num)
	}
	return fmt.Sprint(time.Since(start).Nanoseconds())
}
