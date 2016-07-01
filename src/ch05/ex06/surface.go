//Copyright (c) 2016 by akeboshi. All Rights Reserved.
package main

// copy from ch03/ex01

import (
	"fmt"
	"math"
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
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg'"+
		"style='stroke: grey; fill: white; stroke-width: 0.7'"+
		" width='%d' height ='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, bx, by, cx, cy, dx, dy := getPolygonPoints(i, j)
			if valid(ax, ay, bx, by, cx, cy, dx, dy) {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}

	fmt.Println("</svg>")
}

func getPolygonPoints(i, j int) (float64, float64, float64, float64, float64, float64, float64, float64) {
	ax, ay := corner(i+1, j)
	bx, by := corner(i, j)
	cx, cy := corner(i, j+1)
	dx, dy := corner(i+1, j+1)
	return ax, ay, bx, by, cx, cy, dx, dy
}

func corner(i, j int) (sx, sy float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := hypot(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func valid(val ...float64) bool {
	for _, v := range val {
		if math.IsInf(v, 0) || math.IsNaN(v) {
			return false
		}
	}
	return true
}
func hypot(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
