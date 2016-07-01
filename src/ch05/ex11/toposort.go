//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import "fmt"

var prereqs = map[string][]string{
	"algorithms":     {"data structres"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
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
	var visitAll func(items []string, cyclo []string)

	visitAll = func(items []string, cyclo []string) {
		for _, item := range items {
			for _, c := range cyclo {
				if c == item {
					for _, cc := range cyclo {
						println(cc)
					}
					println("***cycle***")
					break
				}
			}
			if !seen[item] {
				seen[item] = true
				visitAll(m[item], append(cyclo, item))
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	visitAll(keys, []string{})
	return order
}
