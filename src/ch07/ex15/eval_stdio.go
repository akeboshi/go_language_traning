//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch07/ex13/eval"
	"fmt"
	"os"
	"strconv"
)

func main() {
	expr, vars := scanExpr()
	env := scanEnv(vars)
	got := fmt.Sprintf("%.6g", expr.Eval(env))
	fmt.Printf("%s = %s\n", expr, got)
}

func scanExpr() (eval.Expr, map[eval.Var]bool) {
	var in string
	fmt.Println("plz input expr: ")
	fmt.Scan(&in)
	fmt.Printf("\n\n")
	expr, err := eval.Parse(in)
	if err != nil {
		fmt.Fprintln(os.Stderr, err) // parse error
	}

	vars := map[eval.Var]bool{}
	err = expr.Check(vars)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return expr, vars
}

func scanEnv(vars map[eval.Var]bool) eval.Env {
	var in string
	env := eval.Env{}
	for v := range vars {
		for {
			fmt.Printf("please input %s: ", v)
			fmt.Scan(&in)
			n, err := strconv.ParseFloat(in, 64)
			if err == nil {
				env[v] = n
				break
			}
			fmt.Printf("incorrect number: %s\n", in)
		}
	}
	fmt.Println(env)
	return env
}
