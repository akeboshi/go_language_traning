//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"ch07/ex13/eval"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type senv eval.Env

func main() {
	env := senv{}
	mux := http.NewServeMux()
	mux.Handle("/calc", http.HandlerFunc(env.calc))
	mux.Handle("/input", http.HandlerFunc(env.input))
	mux.Handle("/list", http.HandlerFunc(env.list))
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

func checkVars(vars map[eval.Var]bool, env senv) error {
	val := []string{}
	var checked bool
	for v := range vars {
		checked = false
		for e := range env {
			if v == e {
				checked = true
			}
		}
		if !checked {
			val = append(val, string(v))
		}
	}
	if len(val) == 0 {
		return nil
	}
	return fmt.Errorf("need var: %s", strings.Join(val, ","))
}

func (env senv) list(w http.ResponseWriter, req *http.Request) {
	env.print(w)
}

func (env senv) calc(w http.ResponseWriter, req *http.Request) {
	expr, vars := scanExpr(req.URL.Query().Get("expr"))
	err := checkVars(vars, env)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	got := fmt.Sprintf("%.6g", expr.Eval(eval.Env(env)))
	fmt.Fprintf(w, "%s = %s\n", expr, got)
	env.print(w)
}

func (env senv) input(w http.ResponseWriter, req *http.Request) {
	rEnv, err := scanEnv(req.URL.Query())
	for k, v := range rEnv {
		env[k] = v
	}
	if err != nil {
		fmt.Fprintln(w, err)
	}
	env.print(w)
}

func (env senv) print(w http.ResponseWriter) {
	for k, v := range env {
		fmt.Fprintf(w, "%s:%f\n", k, v)
	}
}

func scanExpr(in string) (eval.Expr, map[eval.Var]bool) {
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

func scanEnv(q url.Values) (env senv, err error) {
	in := map[string]string{}
	env = senv{}
	for k, v := range q {
		in[k] = v[0]
	}
	for k, v := range in {
		n, err := strconv.ParseFloat(v, 64)
		if err == nil {
			// k is not checked
			env[eval.Var(k)] = n
		} else {
			err = fmt.Errorf("incorrect number: %s=%s\n", k, v)
		}
	}
	return
}
