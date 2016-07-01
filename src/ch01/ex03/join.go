//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	sample := []string{"foo", "bar", "hoge"}
	measureJoin(joinWithJoinMethod, sample)
	measureJoin(joinWithOperator, sample)
}

func joinWithOperator(words []string) string {
	s, sep := "", ""
	for _, arg := range words {
		s += sep + arg
		sep = " "
	}
	return s
}

func joinWithJoinMethod(words []string) string {
	return strings.Join(words, " ")
}

func measureJoin(f func([]string) string, words []string) {
	start := time.Now()
	for i := 1; i < 1000000; i++ {
		_ = f(words)
	}
	fmt.Println(time.Since(start).Nanoseconds())
}

/* 結果
join:     175869582
operator: 262662182
*/
