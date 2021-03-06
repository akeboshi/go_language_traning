//Copyright (c) 2016 by akeboshi. All Rights Reserved.
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package eval

import "strconv"

// An Expr is an arithmetic expression.
type Expr interface {
	// Eval returns the value of this Expr in the environment env.
	Eval(env Env) float64
	// Check reports errors in this Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
	String() string
}

//!+ast

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'
	x, y Expr
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

type min struct {
	fn   string // min
	args []Expr
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'e', 4, 64)
}

func (v Var) String() string {
	return string(v)
}

func (u unary) String() string {
	return Format(u)
}

func (b binary) String() string {
	return Format(b)
}

func (c call) String() string {
	return Format(c)
}

func (m min) String() string {
	return Format(m)
}

//!-ast
