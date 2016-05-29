package main

import "ch07/ex03/treesort"

func main() {
	foo := []int{5, 3, 2, 6, 4, 8, 10, 9}
	treesort.Sort(foo)
	for _, f := range foo {
		println(f)
	}
}
