package main

import (
	"ch02/ex01/tempconv"
	"ch02/ex02/lengthconv"
	"ch02/ex02/weightconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var input []string
	input = os.Args[1:]
	if len(os.Args[1:]) == 0 {
		var v string
		fmt.Scan(&v)
		input = append(input, v)
	}
	for _, arg := range input {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		showTemple(t)
		showLength(t)
		showWeight(t)
	}
}

func showTemple(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	k := tempconv.Kelvin(t)
	fmt.Println("---- Temple ----")
	fmt.Printf("%s = %s = %s\n", f, tempconv.FToC(f), tempconv.FToK(f))
	fmt.Printf("%s = %s = %s\n", c, tempconv.CToF(c), tempconv.CToK(c))
	fmt.Printf("%s = %s = %s\n", k, tempconv.Kelvin(k), tempconv.KToF(k))
	fmt.Println("----------------")
}

func showLength(t float64) {
	f := lengthconv.Feet(t)
	m := lengthconv.Meter(t)
	fmt.Println("---- Length ----")
	fmt.Printf("%s = %s\n", f, lengthconv.FToM(f))
	fmt.Printf("%s = %s\n", m, lengthconv.MToF(m))
	fmt.Println("----------------")
}

func showWeight(t float64) {
	p := weightconv.Pound(t)
	g := weightconv.Gram(t)
	fmt.Println("---- Weight ----")
	fmt.Printf("%s = %s\n", p, weightconv.PToG(p))
	fmt.Printf("%s = %s\n", g, weightconv.GToP(g))
	fmt.Println("----------------")
}
