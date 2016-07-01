//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"regexp"
	"strings"
)

func main() {
	println(expand("$foo bar hoge", strings.ToUpper))
	println(strings.ToUpper("s string"))
}

func expand(s string, f func(string) string) string {
	reg := regexp.MustCompile(`\$[a-zA-Z]+`)
	return reg.ReplaceAllStringFunc(s, func(s string) string { return f(s[1:]) })
}
