package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms": {"data structres"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":      {"discrete math"},
	"database":             {"data structures"},
	"discrete math":        {"intro to programming"},
	"formal languages":     {"discrete math"},
	"networks":             {"operating systems"},
	"operating systems":    {"data structures", "computer organization"},
	"programming language": {"data structures", "computer organization"},
}

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string][]string)

	visitAll = func(items map[string][]string) {
		for item, mitem := range items {
			if !seen[item] {
				seen[item] = true
				foo := make(map[string][]string)
				for _, bar := range mitem { // mitem = m[item]になってる。
					foo[bar] = m[bar]
				}
				visitAll(foo)
				order = append(order, item)
			}
		}
	}

	visitAll(m)
	return order
}
