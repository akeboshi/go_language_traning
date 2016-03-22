package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	rotate(a[:])
	fmt.Println(a)
}

func rotate(s []int) {
	for i := 0; i < len(s)-1; i = i + 1 {
		n := i - 1
		if n < 0 {
			n = len(s) + n
		}
		s[i], s[n] = s[n], s[i]
	}
}
