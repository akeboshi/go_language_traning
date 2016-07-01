//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	var useSaddle = flag.Bool("s", false, "use saddle method")
	var useEgg = flag.Bool("e", false, "use egg box method")
	var useMogle = flag.Bool("m", false, "use mogle method")
	var useHypot = flag.Bool("h", false, "use hypot method")
	flag.Parse()
	var functions []func(float64, float64) float64
	if *useSaddle {
		functions = append(functions, saddle)
	}
	if *useEgg {
		functions = append(functions, egg)
	}
	if *useMogle {
		functions = append(functions, mogle)
	}
	if *useHypot {
		functions = append(functions, hypot)
	}
	if len(functions) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'"+
		"style='stroke: grey; fill: white; stroke-width: 0.7'"+
		"width='%d' height ='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, bx, by, cx, cy, dx, dy := getPolygonPoints(i, j, functions)
			if valid(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Println("</svg>")
}

func getPolygonPoints(i, j int, functions []func(float64, float64) float64) (float64, float64, float64, float64, float64, float64, float64, float64) {
	ax, ay := corner(i+1, j, functions)
	bx, by := corner(i, j, functions)
	cx, cy := corner(i, j+1, functions)
	dx, dy := corner(i+1, j+1, functions)
	return ax, ay, bx, by, cx, cy, dx, dy
}

func corner(i, j int, functions []func(float64, float64) float64) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := applyMethods(x, y, functions)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func valid(val ...float64) bool {
	for _, v := range val {
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return false
		}
	}
	return true
}

func applyMethods(x, y float64, functions []func(float64, float64) float64) float64 {
	var sums float64
	for _, f := range functions {
		sums += f(x, y)
	}
	functionLen := len(functions)
	if functionLen == 0 {
		functionLen = 1
	}
	return sums / float64(functionLen)
}

func mogle(x, y float64) float64 {
	return math.Sin(-x) * math.Pow(1.5, -math.Hypot(x, y))
}

func egg(x, y float64) float64 {
	return math.Pow(2, math.Sin(y)) * math.Pow(2, math.Sin(x)) / 12
}

func saddle(x, y float64) float64 {
	return math.Sin(x*y/10) / 10
}

func hypot(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
