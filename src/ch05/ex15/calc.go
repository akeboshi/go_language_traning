//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"log"
	"os"
)

var logger = log.New(os.Stderr, "", 11)

func main() {
	println(max(1, 2, 3))
	println(max(3, 2, 1))
	println(notEmptyValsMax(1, 2, 3))
	println(notEmptyValsMax(3, 2, 1))
	println(min(1, 2, 3))
	println(min(3, 2, 1))
	println(notEmptyValsMin(1, 2, 3))
	println(notEmptyValsMin(3, 2, 1))
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(vals ...int) int {
	if len(vals) == 0 {
		logger.Println("no vals")
		os.Exit(1)
	}
	maxVal := vals[0]
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func min(vals ...int) int {
	if len(vals) == 0 {
		logger.Println("no vals")
		os.Exit(1)
	}
	minVal := vals[0]
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func notEmptyValsMax(val0 int, vals ...int) int {
	maxVal := val0
	for _, val := range vals {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func notEmptyValsMin(val0 int, vals ...int) int {
	minVal := val0
	for _, val := range vals {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}
